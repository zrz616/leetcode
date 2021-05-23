package main

import "fmt"

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	p := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		p[i] = target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if p[i] == nums[j] {
				res := []int{i, j}
				return res
			}
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
