# elastic api操作
> https://www.cnblogs.com/buchizicai/p/17093719.html
> https://blog.frognew.com/2021/05/go-elasticsearch-introduction.html

mysql与elasticsearch的概念对比

|Table	|Index|	索引(index)，就是文档的集合，类似数据库的表(table)|
|---|---|---|
|Row	|Document|	文档（Document），就是一条条的数据，类似数据库中的行（Row），文档都是JSON格式|
|Column|	Field|	字段（Field），就是JSON文档中的字段，类似数据库中的列（Column）|
|Schema|	Mapping|	Mapping（映射）是索引中文档的约束，例如字段类型约束。类似数据库的表结构（Schema）|
|SQL	|DSL	|DSL是elasticsearch提供的JSON风格的请求语句，用来操作elasticsearch，实现CRUD|

## 基础
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

## 索引 （表）
- 创建索引库：PUT /索引库名
- 查询索引库：GET /索引库名
- 删除索引库：DELETE /索引库名
- 修改索引库（添加字段）：PUT /索引库名/_mapping

### 列出所有索引以及存储大小
```bash
curl "http://127.0.0.1:9200/_cat/indices?v"
```

### 创建索引
```bash
curl -X PUT "http://127.0.0.1:9200/wwwuser"
```
### 删除索引
```bash
curl -X DELETE "http://127.0.0.1:9200/wwwuser"
```

## 文档操作 （记录）
- 创建文档：POST /{索引库名}/_doc/文档id
- 查询文档：GET /{索引库名}/_doc/文档id
- 删除文档：DELETE /{索引库名}/_doc/文档id
- 修改文档： 
  - 全量修改：PUT /{索引库名}/_doc/文档id
  - 增量修改：POST /{索引库名}/_update/文档id { "doc": {字段}}

### 插入/覆盖文档
```bash
curl —X PUT 'http://127.0.0.1:9200/chentao_demo/_doc/1' \
--header 'Content-Type: application/json' \
--data '{
    "ID":1,
    "username":"alex",
    "age":19
}'
```

### 按ID查询
curl -GET IP:9200/索引名称/索引类型/ID
```bash
 curl "http://127.0.0.1:9200/chentao_demo/_doc/100"
```

### 删除文档
```bash
curl -X DELETE "http://127.0.0.1:9200/chentao_demo/_doc/100"
```

### 修改文档
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

修改age字段
```bash
curl -X POST "localhost:9200/chentao_demo/_update/10001?pretty" -H 'Content-Type: application/json' -d'
{
  "script" : {
    "source": "ctx._source.age += params.count",
    "lang": "painless",
    "params" : {
      "count" : 4
    }
  }
}
'
```

### 检索全部
url IP:9200/索引名称/索引类型/_search
```bash
curl 'http://127.0.0.1:9200/chentao_demo/_search?pretty'
```

### 轻量搜索
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

### 返回指定字段
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