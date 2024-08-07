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
	t1 := uint64(time.Now().Second())
	data := []*types.UserEs{
		{ID: 100, UpdateTime: t1, CreateTime: t1, Nickname: "章泽天"},
		{ID: 101, UpdateTime: t1, CreateTime: t1, Nickname: "张学友"},
	}
	err := m.BatchAdd(context.Background(), data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
