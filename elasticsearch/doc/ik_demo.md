# IK分词器使用

检测ik分词器
```bash
curl -X GET 'http://127.0.0.1:9200/_analyze' \
--header 'Content-Type: application/json' \
--data '{
  "text": "中华人民共和国国歌",
  "analyzer": "ik_smart"
}'
```

## 创建index已经mappings
```bash
curl -X PUT 'http://127.0.0.1:9200/hotel' \
--header 'Content-Type: application/json' \
--data '{
  "mappings": {
    "properties": {
      "id": {
        "type": "keyword"
      },
      "name":{
        "type": "text",
        "analyzer": "ik_max_word",
        "copy_to": "all"
      },
      "address":{
        "type": "keyword",
        "index": false
      },
      "price":{
        "type": "integer"
      },
      "score":{
        "type": "integer"
      },
      "brand":{
        "type": "keyword",
        "copy_to": "all"
      },
      "city":{
        "type": "keyword",
        "copy_to": "all"
      },
      "starName":{
        "type": "keyword"
      },
      "business":{
        "type": "keyword"
      },
      "location":{
        "type": "geo_point"
      },
      "pic":{
        "type": "keyword",
        "index": false
      },
      "all":{
        "type": "text",
        "analyzer": "ik_max_word"
      }
    }
  }
}
'
```