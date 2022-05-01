package _00_最长递增子序列

/*
1、动态规划
2、dp[i]表示包含第i个数的最长严格递增子序列的长度
*/
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	result := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
