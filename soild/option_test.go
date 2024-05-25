package soild

import "testing"

// 执行测试 go test -v -run TestOption
func TestOption(t *testing.T) {
	conn := NewConnection(":8080", WithTimeout(30), WithCache(true))
	if conn.timeout != 30 {
		t.Errorf("Expected 30, but got %d", conn.timeout)
	}
	if conn.cache != true {
		t.Errorf("Expected true, but got %v", conn.cache)
	}
}
