package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0                           // 慢指针，标记唯一元素位置
	for j := 1; j < len(nums); j++ { // 快指针遍历数组
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j] // 将新元素覆盖到唯一区
		}
	}
	return i + 1 // 新数组长度
}

func main() {
	testCases := [][]int{
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		{},
		{5},
	}

	for _, nums := range testCases {
		k := removeDuplicates(nums)
		fmt.Printf("Input: %v → Output: %v (Length: %d)\n", nums, nums[:k], k)
	}
}
