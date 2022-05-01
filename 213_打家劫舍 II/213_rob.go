package _13_打家劫舍_II

/*
1、动态规划
2、滚动数组
3、分区间，避免首尾相接
*/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	res1 := robSub(nums[0 : len(nums)-1])
	res2 := robSub(nums[1:len(nums)])
	return max(res1, res2)
}

func robSub(nums []int) int {
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
