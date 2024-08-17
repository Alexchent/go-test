package test

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"testing"
)

func TestSearch3(t *testing.T) {
	queryAuthor := `from chentao_demo
    | where nick_name == "张学友"
    | sort age desc
    | limit 10`

	response, err := typedClient.Esql.Query().
		Query(queryAuthor).
		Format("csv").
		Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(bytes.NewReader(response))
	rows, err := reader.ReadAll()
	for _, row := range rows {
		fmt.Println(row)
	}
}
