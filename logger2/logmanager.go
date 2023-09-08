package logger

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type OutputEntry struct {
	index  int
	output Output
}

type LogManager struct {
	loggerMap     map[string]*wrappedLogger
	defaultLogger *wrappedLogger
}

func newLogManager() *LogManager {
	return &LogManager{
		loggerMap: make(map[string]*wrappedLogger),
	}
}

func defaultValue(v int, d int) int {
	if v == 0 {
		return d
	}
	return v
}

func buildWriters(outputs []Output) []zapcore.WriteSyncer {
	writers := make([]zapcore.WriteSyncer, 0)
	for _, output := range outputs {
		switch output.GetType() {
		case OutputTypeConsole:
			writers = append(writers, zapcore.AddSync(os.Stdout))
		case OutputTypeFile:
			fileOutput := output.(*FileOutput)
			writers = append(writers, zapcore.AddSync(&lumberjack.Logger{
				LocalTime:  true,
				Filename:   fileOutput.FileName,
				MaxSize:    defaultValue(int(fileOutput.MaxSize), 100),
				MaxBackups: defaultValue(fileOutput.MaxBackups, 30),
				MaxAge:     defaultValue(fileOutput.MaxDays, 15),
			}))
		}
	}
	return writers
}

func buildOutputs(config *Config) []Output {
	if len(config.Outputs) == 0 {
		if len(config.LogDir) == 0 {
			return []Output{&ConsoleOutput{OutputBase: OutputBase{Type: OutputTypeConsole}, Colored: true}}
		} else {
			_, programName := filepath.Split(os.Args[0])
			fileName := filepath.Join(config.LogDir, programName+".log")
			return []Output{&FileOutput{OutputBase: OutputBase{Type: OutputTypeFile}, FileName: fileName}}
		}
	}
	outputs := make([]Output, 0)
	for _, o := range config.Outputs {
		var output Output
		switch o["type"] {
		case OutputTypeConsole:
			output = &ConsoleOutput{}
		case OutputTypeFile:
			output = &FileOutput{}
		default:
			panic(fmt.Errorf("read logger config file error: unknown output type: %s", o["type"]))
		}
		err := decodeOutput(o, output)
		if err != nil {
			panic(fmt.Errorf("read logger config file error: %s", err))
		}
		outputs = append(outputs, output)
		if fileOutput, ok := output.(*FileOutput); ok {
			if !filepath.IsAbs(fileOutput.FileName) {
				if len(config.LogDir) > 0 {
					fileOutput.FileName = filepath.Join(config.LogDir, fileOutput.FileName)
				}
			}
		}
	}
	return outputs
}

func decodeOutput(input map[string]interface{}, output Output) error {
	metaData := &mapstructure.Metadata{}
	err := mapstructure.DecodeMetadata(input, output, metaData)
	if err != nil {
		return err
	}
	if len(metaData.Unused) > 0 {
		return errors.New("unknown field " + strings.Join(metaData.Unused, ","))
	}
	return nil
}

func buildLogger(format string, outputEntries []OutputEntry, writers []zapcore.WriteSyncer) *wrappedLogger {
	ws := make([]zapcore.WriteSyncer, 0, len(outputEntries))
	colored := false

	for _, entry := range outputEntries {
		if entry.output.GetType() == OutputTypeConsole {
			colored = entry.output.(*ConsoleOutput).Colored
		}
		ws = append(ws, writers[entry.index])
	}
	w := zapcore.NewMultiWriteSyncer(ws...)
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if colored {
		encodeConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}
	encodeConfig.ConsoleSeparator = " "
	encoder := zapcore.NewConsoleEncoder(encodeConfig)
	level := zap.NewAtomicLevelAt(zap.InfoLevel)
	core := zapcore.NewCore(encoder, w, level)
	l := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return &wrappedLogger{
		sugured: l.Sugar(),
		level:   level,
	}
}

func (lm *LogManager) LoadWith(config *Config) {
	lm.Reset()
	outputs := buildOutputs(config)
	categoryMap := make(map[string][]OutputEntry)
	var firstCategory string
	for i, output := range outputs {
		for _, category := range output.GetCategories() {
			if len(firstCategory) == 0 {
				firstCategory = category
			}
			if outputItems, ok := categoryMap[category]; ok {
				categoryMap[category] = append(outputItems, OutputEntry{index: i, output: output})
			} else {
				categoryMap[category] = []OutputEntry{{index: i, output: output}}
			}
		}
	}

	writers := buildWriters(outputs)
	for category, outputEntries := range categoryMap {
		lm.loggerMap[category] = buildLogger(config.Format, outputEntries, writers)
	}
	lm.SetLevelByName(config.Level)
	for category, lvlName := range config.Levels {
		lm.SetCategoryLevelByName(category, lvlName)
	}
	lm.defaultLogger = lm.loggerMap[DefaultCategoryName]
	if lm.defaultLogger == nil {
		lm.defaultLogger = lm.loggerMap[firstCategory]
	}
}

func (lm *LogManager) Get(category string) Logger {
	if l, ok := lm.loggerMap[category]; ok {
		return l
	}
	return lm.defaultLogger
}

func (lm *LogManager) SetLevel(lvl Level) {
	for _, l := range lm.loggerMap {
		l.SetLevel(lvl)
	}
}

func (lm *LogManager) SetLevelByName(lvlName string) {
	if len(lvlName) == 0 {
		return
	}
	var l zapcore.Level
	err := l.Set(lvlName)
	if err == nil {
		lm.SetLevel(Level(l))
	}
}

func (lm *LogManager) SetCategoryLevel(category string, lvl Level) {
	if l, ok := lm.loggerMap[category]; ok {
		l.SetLevel(lvl)
	}
}

func (lm *LogManager) SetCategoryLevelByName(category string, lvlName string) {
	if len(lvlName) == 0 {
		return
	}
	var l zapcore.Level
	err := l.Set(lvlName)
	if err == nil {
		lm.SetCategoryLevel(category, Level(l))
	}
}

func (lm *LogManager) Reset() {
	lm.loggerMap = make(map[string]*wrappedLogger)
	lm.defaultLogger = nil
}

func (lm *LogManager) Close() {
	for _, l := range lm.loggerMap {
		l.sugured.Sync()
	}
}
