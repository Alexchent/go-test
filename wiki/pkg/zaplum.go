package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

//go:embed test-data.json
var testdata string

func main() {
	EncoderConfig := zapcore.EncoderConfig{
		MessageKey:          "msg",
		LevelKey:            "",
		TimeKey:             "",
		NameKey:             "",
		CallerKey:           "", // 不记录日志调用位置
		FunctionKey:         zapcore.OmitKey,
		StacktraceKey:       "",
		SkipLineEnding:      false,
		LineEnding:          zapcore.DefaultLineEnding,
		EncodeLevel:         zapcore.LowercaseLevelEncoder,
		EncodeTime:          zapcore.RFC3339TimeEncoder,
		EncodeDuration:      zapcore.SecondsDurationEncoder,
		EncodeCaller:        zapcore.ShortCallerEncoder,
		EncodeName:          nil,
		NewReflectedEncoder: nil,
		ConsoleSeparator:    "",
	}
	// lumberjack.Logger is already safe for concurrent use, so we don't need to
	// lock it.
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "/Users/chentao/go/src/go-test/log/ad.log",
		MaxSize:    5, // megabytes
		MaxBackups: 300,
		MaxAge:     28, // days
	})
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(EncoderConfig),
		w,
		zap.InfoLevel,
	)
	logger := zap.New(core)

	fmt.Println(testdata)
	var data interface{}
	json.Unmarshal([]byte(testdata), &data)
	logger.Info("log", zap.Any("content", data))

	//for i := 0; i < 10000; i++ {
	//	logger.Info(fmt.Sprintf(testdata, i))
	//}
}
