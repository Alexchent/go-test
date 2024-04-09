package main

import (
	"github.com/alexchen/go_test/suanfa/code"
	"github.com/alexchen/go_test/suanfa/solution"
	"github.com/alexchen/go_test/util"
	"go.uber.org/goleak"
	"reflect"
	"testing"
)

// 执行单个测试用例 go test -v -test.run TestAddCommas
// go test -v -run TestAddCommas
func TestAddCommas(t *testing.T) {
	if p := util.AddCommas(100000); p != "100,000" {
		t.Errorf("no pass 1 but %+v got", p)
	}

	if p := util.AddCommas(198); p != "198" {
		t.Errorf("no pass 2 but %+v got", p)
	}

	if p := util.AddCommas(1198); p != "1,198" {
		t.Errorf("no pass 3 but %+v got", p)
	}

	if p := util.AddCommas(-1198); p != "-1,198" {
		t.Errorf("no pass 4 but %+v got", p)
	}
}

func TestBuildCsv(t *testing.T) {
	m := make(map[int][]string)
	m[0] = []string{"学生编号", "学生姓名", "学生特长"}
	m[1] = []string{"s1", "学生1", "乾坤大挪移"}
	m[2] = []string{"s2", "学生2", "乾坤大挪移"}
	m[3] = []string{"s3", "学生3", "乾坤大挪移"}
	m[4] = []string{"s4", "学生4", "乾坤大挪移"}
	m[5] = []string{"s5", "学生5", "乾坤大挪移"}
	m[6] = []string{"s6", "学生6", "乾坤大挪移"}
	m[7] = []string{"s7", "学生7", "乾坤大挪移"}
	m[8] = []string{"s8", "学生8", "乾坤大挪移"}
	m[9] = []string{"s9", "学生9", "乾坤大挪移"}
	m[10] = []string{"s10", "学生10", "乾坤大挪移"}

	//BuildCsvFile(m, "/Users/chentao/Downloads/report/report-data")
	util.BuildCsvFile(m, "report-data")
}

// 测试slice是否存在某个元素
func TestInSlice(t *testing.T) {
	slice := []string{"php", "java", "go", "python"}
	if !util.InSlice(slice, "php") {
		t.Errorf("no pass 1 but %+v got", "php")
	}

	if !util.InSlice(slice, "go") {
		t.Errorf("no pass 2 but %+v got", "go")
	}

	if util.InSlice(slice, "c++") {
		t.Errorf("no pass 3 but %+v got", "c++")
	}
}

// 测试反转slice
func TestReverseSlice(t *testing.T) {
	slice := []string{"php", "java", "go", "python"}
	util.ReverseSlice(slice)
	if slice[0] != "python" {
		t.Errorf("no pass 1 but %+v got", slice[0])
	}

	if slice[1] != "go" {
		t.Errorf("no pass 2 but %+v got", slice[1])
	}

	if slice[2] != "java" {
		t.Errorf("no pass 3 but %+v got", slice[2])
	}

	if slice[3] != "php" {
		t.Errorf("no pass 4 but %+v got", slice[3])
	}
}

func TestDecimalToBinary(t *testing.T) {
	n := 1041
	if p := solution.DecimalToBinary(n); p != "10000010001" {
		t.Errorf("no pass 1 but %+v got", p)
	}
}

func TestFindMissing2(t *testing.T) {
	A := []int{1, 3, 6, 4, 1, 2}
	if p := solution.FindMissing2(A); p != 5 {
		t.Errorf("no pass 1 but %+v got", p)
	}

	A = []int{1, 2, 3}
	if p := solution.FindMissing2(A); p != 4 {
		t.Errorf("no pass 2 but %+v got", p)
	}

	A = []int{-1, -3}
	if p := solution.FindMissing2(A); p != 1 {
		t.Errorf("no pass 3 but %+v got", p)
	}
}

func TestFindEle(t *testing.T) {
	n := []int{}
	if p := code.FindEle(n); p != 3 {
		t.Errorf("no pass 1 but %+v got", p)
	}
}

func TestSolution(t *testing.T) {
	if p := solution.Solution([]int{2, -2, 3, 0, 4, -7}); p != 4 {
		t.Errorf("no pass 1 but %+v got", p)
	}
}

func TestSolution2(t *testing.T) {
	if p := solution.Solution2(2, []int{1}, []int{2}); p != 1 {
		t.Errorf("no pass 1 but %+v got", p)
	}
	if p := solution.Solution2(2, []int{1}, []int{2}); p != 1 {
		t.Errorf("no pass 1 but %+v got", p)
	}
	if p := solution.Solution2(2, []int{1}, []int{2}); p != 1 {
		t.Errorf("no pass 1 but %+v got", p)
	}
	if p := solution.Solution2(2, []int{1}, []int{2}); p != 1 {
		t.Errorf("no pass 1 but %+v got", p)
	}
}

func TestLeak(t *testing.T) {
	defer goleak.VerifyNone(t)

	ch := make(chan int)
	go func() {
		<-ch
	}()
}

func TestSumTwo(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := code.TwoSum(a, 7)
	if !reflect.DeepEqual(b, []int{2, 3}) {
		t.Errorf("no pass 1 but")
	}
}
