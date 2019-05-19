package main

import "fmt"

func twoSum(nums []int, target int) []int {
	ref := make(map[int]int)
	for i, v := range nums {
		if j, ok := ref[target-v]; ok {
			return []int{j, i}
		}
		ref[v] = i
	}
	return []int{-1, -1}
}

func main() {
	result := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(result)
}
