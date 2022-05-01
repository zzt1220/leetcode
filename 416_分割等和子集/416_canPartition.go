package _16_分割等和子集

/*
1、动态规划，01背包思想
2、滚动数组
*/
func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return target == dp[target]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
