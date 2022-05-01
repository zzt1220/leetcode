package _18_零钱兑换_II

/*
1、动态规划，完全背包问题
*/
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = dp[j] + dp[j-coins[i]]
		}
	}
	return dp[amount]
}
