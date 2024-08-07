## 索引

1. 列出所有索引以及存储大小
```bash
curl "http://127.0.0.1:9200/_cat/indices?v"
```
2. 创建索引
```bash
curl -X PUT "http://127.0.0.1:9200/wwwuser"
```

3. 创建mapping
```bash
curl -X PUT "http://127.0.0.1:9200/wwwuser" \
--header 'Content-Type: application/json' \
--data '{
  "mappings": {
    "my_type": {
      "properties": {
        "text": {
          "type": "text",
          "analyzer": "ik_max_word"
        }
      }
    }
  }
}'

```

4. 删除索引
```bash
curl -X DELETE "http://127.0.0.1:9200/wwwuser"
```