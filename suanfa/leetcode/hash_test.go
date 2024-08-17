package main

import (
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "t1", args: args{gas: []int{1, 2, 3, 4, 5}, cost: []int{3, 4, 5, 1, 2}}, want: 3},
		{name: "t2", args: args{gas: []int{2, 3, 4}, cost: []int{3, 4, 3}}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanCompleteCircuit(tt.args.gas, tt.args.cost); got != tt.want {
				t.Errorf("CanCompleteCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCandy(t *testing.T) {
	type args []int
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "t1", args: []int{1, 0, 2}, want: 5},
		{name: "t1", args: []int{1, 3, 4, 5, 2}, want: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Candy(tt.args); got != tt.want {
				t.Errorf("CanCompleteCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}
