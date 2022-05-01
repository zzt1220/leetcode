package _79_完全平方数

/*
1、动态规划，完全背包问题
2、注意初始化，初始化值要取大值
*/
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = i
	}
	for i := 1; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}
	return dp[n]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
