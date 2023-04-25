package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func main() {
	result1 := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(result1)

	result2 := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(result2)

	result3 := twoSum([]int{3, 3}, 6)
	fmt.Println(result3)
}
