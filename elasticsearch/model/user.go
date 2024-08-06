package model

import (
	"context"
	"fmt"
	"github.com/alexchen/go_test/elasticsearch/types"
	"github.com/olivere/elastic/v7"
	"strconv"
	"time"
)

type UserES struct {
	client  *elastic.Client
	index   string
	mapping string
}

var mappingTpl = `{
	"mappings":{
		"properties":{
			"id": 				{ "type": "long" },
			"username": 		{ "type": "keyword" },
			"nickname":			{ "type": "text" },
			"phone":			{ "type": "keyword" },
			"age":				{ "type": "long" },
			"ancestral":		{ "type": "text" },
			"identity":         { "type": "text" },
			"update_time":		{ "type": "long" },
			"create_time":		{ "type": "long" }
			}
		}
	}`
var (
	author  = "chentao"
	project = "demo"
)

func NewUserES(client *elastic.Client) *UserES {
	index := fmt.Sprintf("%s_%s", author, project)
	userEs := &UserES{
		client:  client,
		index:   index,
		mapping: mappingTpl,
	}

	userEs.init()

	return userEs
}

func (es *UserES) init() {
	ctx := context.Background()

	exists, err := es.client.IndexExists(es.index).Do(ctx)
	if err != nil {
		fmt.Printf("userEs init exist failed err is %s\n", err)
		return
	}

	if !exists {
		_, err := es.client.CreateIndex(es.index).Body(es.mapping).Do(ctx)
		if err != nil {
			fmt.Printf("userEs init failed err is %s\n", err)
			return
		}
	}
}

const esRetryLimit = 3

func (es *UserES) BatchAdd(ctx context.Context, user []*types.UserEs) error {
	var err error
	for i := 0; i < esRetryLimit; i++ {
		if err = es.batchAdd(ctx, user); err != nil {
			fmt.Println("batch add failed ", err)
			continue
		}
		return err
	}
	return err
}

func (es *UserES) batchAdd(ctx context.Context, user []*types.UserEs) error {
	req := es.client.Bulk().Index(es.index)
	for _, u := range user {
		u.UpdateTime = uint64(time.Now().UnixNano()) / uint64(time.Millisecond)
		u.CreateTime = uint64(time.Now().UnixNano()) / uint64(time.Millisecond)
		doc := elastic.NewBulkIndexRequest().Id(strconv.FormatUint(u.ID, 10)).Doc(u)
		req.Add(doc)
	}
	if req.NumberOfActions() < 0 {
		return nil
	}
	if _, err := req.Do(ctx); err != nil {
		return err
	}
	return nil
}
