package _6_合并区间

/*
1、贪心算法
2、先按照区间第一个值从小到大排序
3、注意i--不能漏了
*/

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] >= intervals[i+1][0] {
			intervals[i][1] = max(intervals[i][1], intervals[i+1][1])
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			i--
		}
	}
	return intervals
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
