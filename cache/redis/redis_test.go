package redis

import (
	"testing"
	"time"
)

// 在包目录下运行 go 工具，也可以忽略包路径
// % go test
// PASS
// ok      github.com/alexchen/test/cache  0.139s
func TestSet(t *testing.T) {
	Set("test", "hello", time.Second*300)
	v := Get("test")
	if v != "hello" {
		t.Error("失败")
	}
}

func TestStoreMap(t *testing.T) {
	session := map[string]string{"name": "John", "surname": "Smith", "company": "Redis", "age": "29"}
	StoreMap("user-session:123", session)
}

func TestGetMap(t *testing.T) {
	session := GetMap("user-session:123")
	t.Log(session)
	if session["name"] != "John" {
		t.Error("失败")
	}
}
