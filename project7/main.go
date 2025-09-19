package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	// 按区间起始点升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]} // 初始化结果集

	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		last := merged[len(merged)-1]

		if current[0] <= last[1] { // 重叠或相邻
			// 合并区间
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else { // 不重叠
			merged = append(merged, current)
		}
	}

	return merged
}

func main() {
	testCases := []struct {
		input    [][]int
		expected [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}}, // 基础合并
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},                                       // 相邻合并
		{[][]int{}, [][]int{}},                                     // 空输入
		{[][]int{{1, 4}, {2, 3}}, [][]int{{1, 4}}},                 // 完全包含
		{[][]int{{5, 6}, {1, 3}, {3, 4}, {2, 5}}, [][]int{{1, 6}}}, // 多区间合并
	}

	for _, tc := range testCases {
		result := merge(tc.input)
		fmt.Printf("Input: %v → Output: %v (Expected: %v)\n", tc.input, result, tc.expected)
	}
}
