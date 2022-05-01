package _09_斐波那契数

/*
1、动态规划
*/

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var res int
	var dp [2]int
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		res = dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = res
	}
	return res
}
