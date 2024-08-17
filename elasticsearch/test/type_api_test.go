package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"testing"
)

var typedClient *elasticsearch.TypedClient
var tindex = "my-index"

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func init() {
	typedClient = NewOfficeEsClient8()
}

func NewOfficeEsClient8() *elasticsearch.TypedClient {
	client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return client
}

func TestCreateIndex2(t *testing.T) {
	//do, err := typedClient.Indices.Create(tindex).Do(context.TODO())
	//if err != nil {
	//	return
	//}
	//t.Log(do.Index)

	res, err := typedClient.Indices.Create("test-index").
		Request(&create.Request{
			Mappings: &types.TypeMapping{
				Properties: map[string]types.Property{
					"price": types.NewIntegerNumberProperty(),
					"name":  types.NewKeywordProperty(),
					"all":   types.NewTextProperty(),
				},
			},
		}).
		Do(nil)
	if err != nil {
		t.Log(err)
	}
	t.Log(res)
}

func TestIndexingDocument(t *testing.T) {
	id := "1"
	document := struct {
		Name string `json:"name"`
		Age  int64  `json:"age"`
	}{
		"go-elasticsearch",
		18,
	}
	do, err := typedClient.Index(tindex).
		Id(id).
		Request(document).
		Do(context.TODO())
	if err != nil {
		return
	}
	t.Log(do)

	response, err := typedClient.Get(tindex, id).Do(context.TODO())
	if err != nil {
		return
	}
	t.Log(response)
}

func TestGet(t *testing.T) {
	id := "1"
	response, err := typedClient.Get(tindex, id).Do(context.TODO())
	if err != nil {
		return
	}
	if response.Found {
		t.Log(string(response.Source_))
		var res User
		json.Unmarshal(response.Source_, &res)
		t.Log(res.Age, res.Name)
	}
}

func TestSearch2(t *testing.T) {
	//do, err := typedClient.Search().Index(tindex).Do(context.TODO())
	//if err != nil {
	//	return
	//}

	response, err := typedClient.Search().
		Index(tindex).
		Request(&search.Request{
			Query: &types.Query{MatchAll: &types.MatchAllQuery{}},
		}).
		Do(context.TODO())
	if err != nil {
		return
	}
	fmt.Println(response.Hits)
}
