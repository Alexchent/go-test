package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"go.uber.org/zap/zapcore"
)

type Options struct {
	LogConf string // 配置文件路径
	LogDir  string // 输出日志目录
}

type Level zapcore.Level

const (
	DebugLevel Level = Level(zapcore.DebugLevel)
	InfoLevel  Level = Level(zapcore.InfoLevel)
	WarnLevel  Level = Level(zapcore.WarnLevel)
	ErrorLevel Level = Level(zapcore.ErrorLevel)
	FatalLevel Level = Level(zapcore.FatalLevel)
)

type Logger interface {
	SetLevel(lvl Level)
	SetLevelByName(lvlName string)

	// 返回增加额外堆栈忽略层级的Logger
	WithAddCallerSkip(skip int) Logger

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

var defaultLogger Logger
var logManager = newLogManager()

type Option interface {
	apply(*Options)
}

type optionFunc func(*Options)

func (f optionFunc) apply(opts *Options) {
	f(opts)
}

func init() {
	// 默认初始化console日志输出，这样当应用层未初始化logger时，也能正常使用logger的功能
	// 当应用层调用Init或InitWithConfig时，会重新初始化
	InitWithConfig(&Config{})
}

func WithConfFile(fileName string) Option {
	return optionFunc(func(options *Options) {
		options.LogConf = fileName
	})
}

func WithConfDir(dir string) Option {
	return optionFunc(func(options *Options) {
		options.LogConf = filepath.Join(dir, DefaultConfigFileName)
	})
}

func WithLogDir(dir string) Option {
	return optionFunc(func(options *Options) {
		options.LogDir = dir
	})
}

// Init 初始化logger
// 如果logConf为空，则使用控制台方式输出
// 如果logConf路径不存在，则panic
// 如果logConf为目录，则使用DefaultConfigFileName定义的默认文件名
// 配置文件中logDir允许被WithLogDir设定的参数覆盖
func Init(options ...Option) {
	opts := &Options{}
	for _, option := range options {
		option.apply(opts)
	}
	logConf := opts.LogConf
	if len(logConf) == 0 { // 如果没有配置文件，则根据logDir是否配置来判断是否记录文件
		InitWithConfig(&Config{LogDir: opts.LogDir})
		return
	}
	info, err := os.Stat(logConf)
	if !os.IsNotExist(err) && info.IsDir() {
		logConf = filepath.Join(logConf, DefaultConfigFileName)
	}

	var config Config
	rb, err := ioutil.ReadFile(logConf)
	if err == nil {
		err = toml.Unmarshal([]byte(replaceEnvVars(string(rb))), &config)
	}
	if err != nil {
		panic(fmt.Errorf("read log config file error: %s", err))
	}
	if len(opts.LogDir) > 0 { // 传入的logDir配置优先于配置文件中的logDir
		config.LogDir = opts.LogDir
	}
	InitWithConfig(&config)
}

func InitWithConfig(config *Config) {
	logManager.LoadWith(config)
	defaultLogger = logManager.defaultLogger.WithAddCallerSkip(1).(*wrappedLogger)
}

func replaceEnvVars(content string) string {
	_, programName := filepath.Split(os.Args[0])
	content = strings.ReplaceAll(content, EnvProgramName.Ref(), programName)
	return content
}

func SetLevel(lvl Level) {
	logManager.SetLevel(lvl)
}

func SetLevelByName(lvlName string) {
	logManager.SetLevelByName(lvlName)
}

func SetCategoryLevel(category string, lvl Level) {
	logManager.SetCategoryLevel(category, lvl)
}

func SetCategoryLevelByName(category string, lvlName string) {
	logManager.SetCategoryLevelByName(category, lvlName)
}

// Get 获取category分类的Logger
func Get(category string) Logger {
	return logManager.Get(category)
}

// Close 结束logger
func Close() {
	logManager.Close()
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}
