package _09_长度最小的子数组

import "math"

/*
1、滑动窗口
*/
func minSubArrayLen(target int, nums []int) int {
	result := math.MaxInt32
	i := 0
	sum := 0
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			subLen := j - i + 1
			if subLen < result {
				result = subLen
			}
			sum -= nums[i]
			i++
		}
	}
	if result == math.MaxInt32 {
		return 0
	}
	return result
}
