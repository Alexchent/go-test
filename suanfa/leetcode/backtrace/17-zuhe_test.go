package hs

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "t1", args: args{
			candidates: []int{2, 3, 6, 7},
			target:     7,
		}, want: [][]int{{2, 2, 3}, {7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateParenthesis(t *testing.T) {
	res := generateParenthesis(3)
	fmt.Println(res)
}

func Test_exist(t *testing.T) {
	b := [][]string{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}}
	r := exist(b, "ABCCED")
	t.Log(r)
}
