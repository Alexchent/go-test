package soild

import (
	"testing"
)

// 执行测试 go test -v -run TestAddStrategy
func TestAddStrategy(t *testing.T) {
	add := &Add{}
	operation := &Operation{}
	operation.SetStrategy(add)

	result := operation.ExecuteStrategy(5, 3)
	if result != 8 {
		t.Errorf("Expected 8, but got %d", result)
	}

	result = operation.ExecuteStrategy(-5, -3)
	if result != -8 {
		t.Errorf("Expected -8, but got %d", result)
	}
}

func TestSubStrategy(t *testing.T) {
	sub := &Sub{}
	operation := &Operation{}
	operation.SetStrategy(sub)

	result := operation.ExecuteStrategy(5, 3)
	if result != 2 {
		t.Errorf("Expected 2, but got %d", result)
	}

	result = operation.ExecuteStrategy(-5, -3)
	if result != -2 {
		t.Errorf("Expected -2, but got %d", result)
	}
}

func TestStrategySwitch(t *testing.T) {
	add := &Add{}
	sub := &Sub{}
	operation := &Operation{}

	operation.SetStrategy(add)
	result := operation.ExecuteStrategy(5, 3)
	if result != 8 {
		t.Errorf("Expected 8, but got %d", result)
	}

	operation.SetStrategy(sub)
	result = operation.ExecuteStrategy(5, 3)
	if result != 2 {
		t.Errorf("Expected 2, but got %d", result)
	}
}
