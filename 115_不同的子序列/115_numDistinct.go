package _15_不同的子序列

/*
1、动态规划，编辑距离问题
2、当s[i]与t[j]相等时，dp[i][j] = dp[i-1][j-1]+dp[i-1][j]
3、本题难度较大
*/
func numDistinct(s string, t string) int {
	if s == "" || t == "" {
		return 0
	}

	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(t))
	}

	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == t[0] {
			count++
		}
		dp[i][0] = count
	}
	for i := 1; i < len(s); i++ {
		for j := 1; j < len(t); j++ {
			if s[i] == t[j] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(s)-1][len(t)-1]
}
