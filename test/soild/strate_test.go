package main

import (
	"github.com/alexchen/go_test/SOILD"
	"testing"
)

// 执行测试 go test -v -run TestAddStrategy
func TestAddStrategy(t *testing.T) {
	add := &SOILD.Add{}
	operation := &SOILD.Operation{}
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
	sub := &SOILD.Sub{}
	operation := &SOILD.Operation{}
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
	add := &SOILD.Add{}
	sub := &SOILD.Sub{}
	operation := &SOILD.Operation{}

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
