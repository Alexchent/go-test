package main

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"os"

	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// 创建一个LumberJack实例，用于日志文件的切割和归档
	lumberJack := &lumberjack.Logger{
		Filename:   "app.log", // 日志文件名
		MaxSize:    10,        // 每个日志文件的最大大小（MB）
		MaxBackups: 3,         // 最大保留的日志文件数
		MaxAge:     28,        // 日志文件保留的最大天数
		Compress:   true,      // 是否压缩归档文件
	}

	// 创建一个Zap实例，用于日志记录
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 设置时间格式
	encoder := zapcore.NewJSONEncoder(encoderConfig)      // 使用JSON格式编码器
	core := zapcore.NewCore(encoder, zapcore.AddSync(lumberJack), zapcore.InfoLevel)
	logger := zap.New(core)

	// 记录日志
	logger.Info("这是一条信息日志")
	logger.Warn("这是一条警告日志")
	logger.Error("这是一条错误日志")

	// 关闭LumberJack实例
	if err := lumberJack.Close(); err != nil {
		fmt.Printf("无法关闭LumberJack实例：%v", err)
		os.Exit(1)
	}
}
