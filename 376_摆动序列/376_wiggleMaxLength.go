package _76_摆动序列

/*
1、贪心算法，计算波峰和波谷数
*/

func wiggleMaxLength(nums []int) int {
	curDiff := 0
	preDiff := 0
	result := 1
	for i := 0; i < len(nums)-1; i++ {
		curDiff = nums[i+1] - nums[i]
		if preDiff <= 0 && curDiff > 0 || preDiff >= 0 && curDiff < 0 {
			result++
			preDiff = curDiff
		}
	}
	return result
}
