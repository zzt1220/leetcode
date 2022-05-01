package _98_打家劫舍

/*
1、动态规划
2、滚动数组
*/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, 2)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		tmp := max(dp[1], dp[0]+nums[i])
		dp[0] = dp[1]
		dp[1] = tmp
	}
	return dp[1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
