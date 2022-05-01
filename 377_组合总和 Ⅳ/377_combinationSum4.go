package _77_组合总和__

/*
1、动态规划，完全背包，这里是求排列总数，注意和组合总数的区别，可以与第510道题比对着分析
2、排列总数和组合总数区别：
如果求组合数就是外层for循环遍历物品，内层for遍历背包
如果求排列数就是外层for遍历背包，内层for循环遍历物品
*/
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if i-nums[j] >= 0 {
				dp[i] = dp[i] + dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}
