sql查询
format支持：txt、csv、json
```bash
curl -X POST "localhost:9200/_sql?format=txt&pretty" -H 'Content-Type: application/json' -d'
    {
      "query": "SELECT * FROM chentao_demo WHERE match(nickname, \u0027张学友\u0027) ORDER BY age DESC LIMIT 5"
    }
    '
```

```bash
curl -X POST "localhost:9200/_sql?format=txt&pretty" -H 'Content-Type: application/json' -d'
{
  "query": "SELECT * FROM chentao_demo WHERE age = 19" 
}
'
```