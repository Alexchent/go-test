package model

import (
	"fmt"
	"github.com/olivere/elastic/v7" // Deprecated: Use the official Elasticsearch client for Go at
	"log"
	"os"
)

func NewEsClient(conf *Config) *elastic.Client {
	url := fmt.Sprintf("http://%s:%d", conf.Elastic.Host, conf.Elastic.Port)
	client, err := elastic.NewClient(
		//elastic 服务地址
		elastic.SetURL(url),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	if err != nil {
		log.Fatalln("Failed to create elastic client for url=" + url)
	}
	return client
}

type Config struct {
	Elastic struct {
		Host string
		Port int64
	}
}
