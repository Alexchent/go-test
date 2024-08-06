package main

import (
	"context"
	"fmt"
	"github.com/alexchen/go_test/elasticsearch/model"
	"github.com/alexchen/go_test/elasticsearch/types"
	"time"
)

func main() {
	var conf model.Config
	conf.Elastic.Host = "127.0.0.1"
	conf.Elastic.Port = 9200
	m := model.NewUserES(model.NewEsClient(&conf))
	data := []*types.UserEs{
		{ID: 100, UpdateTime: uint64(time.Now().Second()), CreateTime: uint64(time.Now().Second())},
	}
	err := m.BatchAdd(context.Background(), data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
