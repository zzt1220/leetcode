package _0_爬楼梯

/*
1、动态规划
*/
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	var res int
	var dp [2]int
	dp[0] = 1
	dp[1] = 2
	for i := 3; i <= n; i++ {
		res = dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = res
	}
	return res
}
