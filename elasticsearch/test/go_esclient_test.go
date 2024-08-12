package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	ela "github.com/elastic/go-elasticsearch/v7"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

var client *ela.Client
var index = "book-0.1.0"

func init() {
	client = NewOfficeEsClient()
}

func NewOfficeEsClient() *ela.Client {
	client, err := ela.NewClient(ela.Config{
		Addresses: []string{"http://localhost:9200"},
	})
	if err != nil {
		return nil
	}
	return client
}

// go test -v -run TestNewEsClient
func TestNewEsClient(t *testing.T) {
	c := NewOfficeEsClient()
	info, err := c.Info()
	if err != nil {
		return
	}
	fmt.Println(info)
}

func TestCreateIndex(t *testing.T) {
	a := assert.New(t)
	response, err := client.Indices.Create(index, client.Indices.Create.WithBody(strings.NewReader(`
	{
		"aliases": {
			"book":{}
		},
		"settings": {
			"analysis": {
				"normalizer": {
					"lowercase": {
						"type": "custom",
						"char_filter": [],
						"filter": ["lowercase"]
					}
				}
			}
		},
		"mappings": {
			"properties": {
				"name": {
					"type": "keyword",
					"normalizer": "lowercase"
				},
				"price": {
					"type": "double"
				},
				"summary": {
					"type": "text",
					"analyzer": "ik_max_word"
				},
				"author": {
					"type": "keyword"
				},
				"pubDate": {
					"type": "date"
				},
				"pages": {
					"type": "integer"
				}
			}
		}
	}
	`)))
	a.Nil(err)
	t.Log(response)
}

// 同时创建index、mapping、doc
func TestIndexDoc(t *testing.T) {
	document := struct {
		Name string `json:"name"`
		Age  int64  `json:"age"`
	}{
		"go-elasticsearch",
		18,
	}
	data, _ := json.Marshal(document)
	response, err := client.Index(index, bytes.NewReader(data))
	if err != nil {
		return
	}
	t.Log(response)
}

func TestGetIndex(t *testing.T) {
	a := assert.New(t)
	response, err := client.Indices.Get([]string{index})
	a.Nil(err)
	t.Log(response)
}

type Book struct {
	Author  string     `json:"author"`
	Name    string     `json:"name"`
	Pages   int        `json:"pages"`
	Price   float64    `json:"price"`
	PubDate *time.Time `json:"pubDate"`
	Summary string     `json:"summary"`
}

func TestCreateDocument(t *testing.T) {
	a := assert.New(t)
	body := &bytes.Buffer{}
	pubDate := time.Now()
	err := json.NewEncoder(body).Encode(&Book{
		Author:  "金庸",
		Price:   96.0,
		Name:    "天龙八部",
		Pages:   1978,
		PubDate: &pubDate,
		Summary: "...",
	})
	a.Nil(err)
	response, err := client.Create("book-0.1.0", "10001", body)
	a.Nil(err)
	t.Log(response)
}

// 更新文档
func TestIndexDocument(t *testing.T) {
	a := assert.New(t)
	body := &bytes.Buffer{}
	pubDate := time.Now()
	err := json.NewEncoder(body).Encode(&Book{
		Author:  "金庸",
		Price:   96.0,
		Name:    "天龙八部2",
		Pages:   1988,
		PubDate: &pubDate,
		Summary: "...",
	})
	a.Nil(err)
	response, err := client.Index("book-0.1.0", body, client.Index.WithDocumentID("10001"))
	a.Nil(err)
	t.Log(response)
}

type doc struct {
	Doc interface{} `json:"doc"`
}

// 局部更新
func TestPartialUpdateDocument(t *testing.T) {
	a := assert.New(t)
	body := &bytes.Buffer{}
	err := json.NewEncoder(body).Encode(&doc{
		Doc: &Book{
			Name: "《天龙八部》",
		},
	})
	t.Log(body)
	a.Nil(err)
	response, err := client.Update("book-0.1.0", "10001", body)
	a.Nil(err)
	t.Log(response)

	//client.Update("my_index", "id", strings.NewReader(`{doc: { language: "Go" }}`))
}

func TestSearch(t *testing.T) {
	query := `{ "query": { "match_all": {} } }`
	search, err := client.Search(
		client.Search.WithIndex(index),
		client.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		return
	}
	t.Log(search)
}

func TestDeleteDoc(t *testing.T) {
	response, err := client.Delete(index, "10001")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(response)
}

func TestDeleteIndex(t *testing.T) {
	response, err := client.Indices.Delete([]string{index})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(response)
}
