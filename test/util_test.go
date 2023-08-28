package main

import (
	"github.com/alexchen/go_test/util"
	"testing"
)

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
	m[0] = []string{"学生编号", "学生姓名", "学生特长", "序号"}
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
	util.BuildCsvFile(m, "../log/report-data2")
}
