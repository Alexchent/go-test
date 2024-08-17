package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	hashT := make(map[int]int, len(nums))
	for i, v := range nums {
		if p, ok := hashT[target-v]; ok {
			return []int{p, i}
		}
		hashT[v] = i
	}
	return []int{}
}

func groupAnagrams(strs []string) [][]string {
	hashMap := map[string][]string{}
	for _, str := range strs {
		// 排序
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sorted := string(s)
		hashMap[sorted] = append(hashMap[sorted], str)
	}
	ans := make([][]string, 0, len(hashMap))
	for _, v := range hashMap {
		ans = append(ans, v)
	}
	return ans
}

func CanCompleteCircuit(gas []int, cost []int) int {
	totalTank := 0
	currTank := 0
	startStation := 0

	for i := 0; i < len(gas); i++ {
		totalTank += gas[i] - cost[i]
		currTank += gas[i] - cost[i]

		// 如果当前油箱油量小于 0，更新起始加油站为下一个，并重置当前油箱油量
		if currTank < 0 {
			currTank = 0
			startStation = i + 1
		}
	}

	// 如果总油箱油量大于等于 0，返回起始加油站编号，否则返回 -1
	if totalTank >= 0 {
		return startStation
	}
	return -1
}

func Candy(ratings []int) int {
	n := len(ratings)
	left := make([]int, n)
	for i, _ := range ratings {
		left[i] = 1
	}
	fmt.Println(left)
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	fmt.Println(left)
	count := left[n-1]
	for k := n - 2; k >= 0; k-- {
		if ratings[k] > ratings[k+1] {
			left[k] = max(left[k+1]+1, left[k])
		}
		count += left[k]
	}
	fmt.Println(left)
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// ThreeSum 三数之和
// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
// 同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var result [][]int
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			return result
		}
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
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
	return result
}
