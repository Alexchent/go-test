package main

import (
	"github.com/alexchen/go_test/suanfa/code"
	"reflect"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	if v := code.MajorityElement(nums); v != 2 {
		t.Error("error, want 2 got ", v)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	num := []int{1, 1, 1, 2, 3, 3, 4, 5, 12, 20}
	//s := BinarySearch(num, 3)
	//fmt.Println(s)
	l := code.RemoveDuplicates(num)
	//fmt.Println(num[:l])
	mb := []int{1, 2, 3, 4, 5, 12, 20}
	if !reflect.DeepEqual(num[:l], mb) {
		t.Error("error, want 7 got ", l)
	}
}

func TestReverse(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	code.Reverse(nums, 0, len(nums)-1)
	//t.Log(nums)
	if !reflect.DeepEqual(nums, []int{7, 6, 5, 4, 3, 2, 1}) {
		t.Error("error: ", nums)
	}
}

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	code.Rotate(nums, 3)
	//t.Log(nums)
	if !reflect.DeepEqual(nums, []int{5, 6, 7, 1, 2, 3, 4}) {
		t.Error("error: ", nums)
	}
}

func TestJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	if v := code.Jump(nums); v != 2 {
		t.Error("error, want 2 got ", v)
	}
}
