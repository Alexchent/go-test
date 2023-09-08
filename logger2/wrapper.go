package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type wrappedLogger struct {
	sugured *zap.SugaredLogger
	level   zap.AtomicLevel
}

func (l *wrappedLogger) SetLevel(lvl Level) {
	l.level.SetLevel(zapcore.Level(lvl))
}

func (l *wrappedLogger) SetLevelByName(lvlName string) {
	var level zapcore.Level
	err := level.Set(lvlName)
	if err == nil {
		l.level.SetLevel(level)
	}
}

func (l *wrappedLogger) WithAddCallerSkip(skip int) Logger {
	return &wrappedLogger{
		sugured: l.sugured.Desugar().WithOptions(zap.AddCallerSkip(skip)).Sugar(),
		level:   l.level,
	}
}

func (l *wrappedLogger) Debugf(format string, args ...interface{}) {
	l.sugured.Debugf(format, args...)
}

func (l *wrappedLogger) Infof(format string, args ...interface{}) {
	l.sugured.Infof(format, args...)
}

func (l *wrappedLogger) Warnf(format string, args ...interface{}) {
	l.sugured.Warnf(format, args...)
}

func (l *wrappedLogger) Errorf(format string, args ...interface{}) {
	l.sugured.Errorf(format, args...)
}

func (l *wrappedLogger) Fatalf(format string, args ...interface{}) {
	l.sugured.Fatalf(format, args...)
}
