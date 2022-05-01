package _22_零钱兑换

/*
1、动态规划，完全背包问题
2、初始化dp[0] = 0，其它dp[i] = i + 1
3、需要条件dp[j-coins[i]] != j-coins[i]+1，才能递推
4、dp[amount] == amount+1，返回-1
*/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = i + 1
	}
	dp[0] = 0
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != j-coins[i]+1 {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
