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
