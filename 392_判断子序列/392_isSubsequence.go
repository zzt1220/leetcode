package _92_判断子序列

/*
1、动态规划，编辑距离问题
*/
func isSubsequence(s string, t string) bool {
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(t)+1; j++ {
			if s[i-1] == t[j-1] {
				dp[i%2][j] = dp[(i-1)%2][j-1] + 1
			} else {
				dp[i%2][j] = dp[i%2][j-1]
			}
		}
	}
	if dp[len(s)%2][len(t)] == len(s) {
		return true
	}
	return false
}
