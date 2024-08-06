
## 查看elastic版本信息
```bash
curl "http://127.0.0.1:9200/"
```

## 检查elastic是否健康
```bash
curl "http://127.0.0.1:9200/_cat/health?v"
```

## 查看节点列表
```bash
curl "http://127.0.0.1:9200/_cat/nodes?v"
```

## 列出所有索引以及存储大小
```bash
curl "http://127.0.0.1:9200/_cat/indices?v"
```

## 创建索引
```bash
curl -X PUT "http://127.0.0.1:9200/wwwuser"
```

## 创建索引的同时设置分词器和过滤器
```bash
curl --location --request PUT 'http://127.0.0.1:9200/fuser' \
--header 'Content-Type: application/json' \
--data '{
    "settings":{
           "number_of_shards":1,     
           "number_of_replicas":2, 
           "index":{
                  "analysis":{
                         "analyzer":{
                                "default":{
                                       "tokenizer":"standard",   
                                       "filter":[ 
                                              "asciifolding",
                                              "lowercase",
                                              "ourEnglishFilter"
                                       ]
                                }
                         },
                         "filter":{
                                "ourEnglishFilter":{
                                       "type":"kstem"
                                }
                         }
                  }
           }
    }
}'
```


## 搜索文档
url IP:9200/索引名称/索引类型/_search
索引类型可以没有
```bash
curl "http://127.0.0.1:9200/chentao_demo/_search"
```

## 按ID查询
curl -GET IP:9200/索引名称/索引类型/ID
```bash
 curl "http://127.0.0.1:9200/chentao_demo/_doc/100"
```

## 删除文档
```bash
curl -X DELETE "http://127.0.0.1:9200/chentao_demo/_doc/100"
```

## 插入/覆盖文档
```bash
curl --location --request PUT 'http://127.0.0.1:9200/chentao_demo/_doc/1' \
--header 'Content-Type: application/json' \
--data '{
    "ID":1,
    "username":"alex",
    "age":19
}'
```

## 修改文档
```bash
curl 'http://127.0.0.1:9200/chentao_demo/_update/1' \
--header 'Content-Type: application/json' \
--data '{
    "doc":{
        "ID":1,
        "username":"tom",
        "age":20
    }
}'
```

## 检索全部
```bash
curl 'http://127.0.0.1:9200/chentao_demo/_search?pretty'
```

## 轻量搜索
query：告诉我们定义查询
match_all：运行简单类型查询指定索引中的所有文档
size：返回结果集数量
```bash
curl 'http://127.0.0.1:9200/chentao_demo/_search?pretty' \
--header 'Content-Type: application/json' \
--data '{
    "query":{
        "match_all":{}
    },
    "size":2
}'
```

### 分页查询 
from：指定文档索引从哪里开始，默认从0开始
size：从from开始，返回多个文档
sort：排序，排序字段与排序方式
```bash
curl 'http://127.0.0.1:9200/chentao_demo/_search?pretty' \
--header 'Content-Type: application/json' \
--data '{
    "query":{
        "match_all":{}
    },
    "from":1,
    "size":5,
    "sort": {"age": {"order":"desc"}}
}'
```

## 返回指定字段
match     模糊匹配 对应 txt类型
match_all 精确匹配 对应 keywords类型
```bash
curl 'http://127.0.0.1:9200/chentao_demo/_search?pretty' \
--header 'Content-Type: application/json' \
--data '{
    "query":{
        "match":{
          "username":"alex"
        }
    },
    "_source": [
      "username",
      "age"
    ]
}'
```


## 一旦es磁盘出现爆满，所有的index都会被自动设置为read-only，重启es无法解除该状态，需要发送请求，如下
```bash
curl -X PUT 'http://127.0.0.1:9200/_settings' \
--header 'Content-Type: application/json' \
--data '{
    "index": {
        "blocks": {
        "read_only_allow_delete": "false"
        }
    }
}'
```

```bash
curl -X PUT 'http://127.0.0.1:9200/_all/_settings' \
--header 'Content-Type: application/json' \
--data '{
    "index.blocks.read_only_allow_delete": null
}'
```