package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newCustomLogger() (*zap.Logger, error) {
	cfg := zap.Config{
		//Level:             zap.AtomicLevel{},
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:          "message",
			LevelKey:            "level",
			TimeKey:             "time",
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
		},
		OutputPaths:      []string{"stdout", "test.log"},
		ErrorOutputPaths: []string{"error.log"},
		InitialFields:    nil,
	}
	return cfg.Build()
}

func main() {
	logger, _ := newCustomLogger()
	defer logger.Sync()
	logger = logger.WithOptions(zap.AddCallerSkip(100))
	logger.Info("info message")
	logger.Debug("debug message")
	logger.Error("error message")
	// 生产环境
	//{
	//	logger, _ := zap.NewProduction()
	//	defer logger.Sync() // 刷新 buffer，保证日志最终会被输出
	//
	//	url := "https://jianghushinian.cn/"
	//	logger.Info("production failed to fetch URL",
	//		zap.String("url", url), // 因为没有使用 interface{} 和反射机制，所以需要指定具体类型
	//		zap.Int("attempt", 3),
	//		zap.Duration("backoff", time.Second),
	//	)
	//}

	// 开发环境
	//{
	//	logger, _ := zap.NewDevelopment()
	//	defer logger.Sync()
	//
	//	url := "https://jianghushinian.cn/"
	//	logger.Debug("development failed to fetch URL",
	//		zap.String("url", url),
	//		zap.Int("attempt", 3),
	//		zap.Duration("backoff", time.Second),
	//	)
	//}
}
