## 索引

列出所有索引以及存储大小
```bash
curl "http://127.0.0.1:9200/_cat/indices?v"
```

创建索引
```bash
curl -X PUT "http://127.0.0.1:9200/chentao_demo"
```

创建mapping
```bash
curl -X PUT "http://127.0.0.1:9200/chentao_demo" \
--header 'Content-Type: application/json' \
--data '{
  "mappings": {
      "properties": {
        "text": {
          "type": "text",
          "analyzer": "ik_max_word"
        }
      }
  }
}'
```

查看索引mapping
```bash
curl -XGET "http://127.0.0.1:9200/chentao_demo/_mapping?pretty" 
```

删除索引
```bash
curl -X DELETE "http://127.0.0.1:9200/chentao_demo"
```