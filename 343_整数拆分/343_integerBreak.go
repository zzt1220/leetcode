package _43_整数拆分

/*
1、动态规划
分为以下两种情况：
j*(i-j)：拆分为2个数相加
j*dp[i-j])：拆分为3个及以上数相加
*/
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
