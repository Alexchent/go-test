package main

import (
	"sort"
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

func ThreeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			return res
		}
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				if left < right && nums[left] == nums[left+1] {
					left++
				}
				if left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
