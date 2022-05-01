package _35_无重叠区间

/*
1、按照区间右值排序
2、先计算不相交区间最大数量（已经排序好）
*/

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	end := intervals[0][1]
	count := 1
	for i := 1; i < len(intervals); i++ {
		if end <= intervals[i][0] {
			end = intervals[i][1]
			count++
		}
	}
	return len(intervals) - count
}
