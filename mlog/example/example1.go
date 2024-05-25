package main

import (
	"github.com/alexchen/go_test/mlog"
	"log"
	"os"
)

func main() {
	mlog.Info("std log")
	mlog.SetOptions(mlog.WithLevel(mlog.DebugLevel))
	mlog.Debug("change std log to debug level")
	mlog.SetOptions(mlog.WithFormatter(&mlog.JsonFormatter{IgnoreBasicFields: false}))
	mlog.Debug("log in json format")
	mlog.Info("another log in json format")

	// 输出到文件
	fd, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("create file test.log failed")
	}
	defer fd.Close()

	l := mlog.New(mlog.WithLevel(mlog.InfoLevel),
		mlog.WithOutput(fd),
		mlog.WithFormatter(&mlog.JsonFormatter{IgnoreBasicFields: false}),
	)
	l.Info("custom log with json formatter")
}
