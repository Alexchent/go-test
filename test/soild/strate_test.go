package main

import (
	"github.com/alexchen/go_test/soild"
	"testing"
)

// 执行测试 go test -v -run TestAddStrategy
func TestAddStrategy(t *testing.T) {
	add := &soild.Add{}
	operation := &soild.Operation{}
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
	sub := &soild.Sub{}
	operation := &soild.Operation{}
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
	add := &soild.Add{}
	sub := &soild.Sub{}
	operation := &soild.Operation{}

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
