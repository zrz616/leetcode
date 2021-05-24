package main

import "fmt"

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	// p := make([]int, len(nums))
	p := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := p[target-nums[i]]
		if ok {
			res := []int{p[target-nums[i]], i}
			return res
		}
		p[nums[i]] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
