package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path"
	"time"
)

var record *FileLogger

func init() {
	record = NewLogger()
}

func Info(logFile string, info ...interface{}) {
	log := record.Output(logFile)
	log.Info(info)
}

func InfoWithFields(logFile string, field logrus.Fields, info ...interface{}) {
	log := record.Output(logFile)
	log.WithFields(field).Info(info)
}

func Infof(logFile string, format string, info ...interface{}) {
	log := record.Output(logFile)
	log.Infof(format, info...)
}

func Error(logFile string, info ...interface{}) {
	log := record.Output(logFile)
	log.Error(info)
}

type FileLogger struct {
	log *logrus.Logger
}

func NewLogger() *FileLogger {
	l := logrus.New()
	l.SetLevel(logrus.DebugLevel)
	//l.SetReportCaller(true)
	l.SetFormatter(&logrus.JSONFormatter{
		//PrettyPrint: true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return &FileLogger{log: l}
}

func (l *FileLogger) Output(logFile string) *logrus.Logger {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "logs/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}

	logFileName := logFile + "-" + now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}
	//写入文件
	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("create log err", err)
	}

	//设置输出
	l.log.SetOutput(io.MultiWriter(writer1, writer2, writer3))
	//logger.Out = writer3
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}

	return l.log
}
