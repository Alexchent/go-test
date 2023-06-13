package sunfa

import (
	"fmt"
	"reflect"
	"testing"
)

// 单元测试样例
// 测试单个文件 go test -v level_01.go suanfa_test.go -cover
// 测试所用用例 go test -v -cover

func TestFib(t *testing.T) {
	num := 5
	res := []int{0, 1, 1, 2, 3}
	data := Fibonacci(5)
	if len(data) != num {
		t.Errorf("len(data) = %d, want, but give %d", num, len(data))
	}

	if !reflect.DeepEqual(data, res) {
		t.Errorf("res = %#v, want %#v", data, res)
	}

}

func TestSortPop(t *testing.T) {
	a := []int{2, 3, 1, 6, 8, 7, 4, 5}
	want := []int{1, 2, 3, 4, 5, 6, 7, 8}
	SortPop(&a)
	if !reflect.DeepEqual(a, want) {
		t.Errorf("a = %#v, want %#v", a, want)
	}
}

func TestFactorial(t *testing.T) {
	res := Factorial(4)
	fmt.Println(res)
	want := 24
	if res != want {
		t.Errorf("res = %#v, want %#v", res, want)
	}
}
