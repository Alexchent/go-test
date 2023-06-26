# 用protobuf作为接口的传输协议

## 1. 准备proto文件
```protobuf
syntax = "proto3";

option go_package = "pb";

/** 
执行以下脚本生成go文件
protoc --go_out=. --go_opt=paths=source_relative proto/person.proto
*/

message PersonListRequest{
  string name=1;
  int32 age=2;
}

message PersonListResponse{
  repeated Person people = 1;
  int32 count = 2;
}

message Person {
  int32 id =1;
  string name=2;
  int32 age=3;
}
```

2. 生成pb.go文件 `protoc --go_out=. --go_opt=paths=source_relative proto/person.proto`

## 写接口，代码参考 [serve](./wiki/protobuf/serve.go)
## 模拟请求，代码参考 [client](./wiki/protobuf/client.go)