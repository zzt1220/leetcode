package _74_一和零

/*
1、动态规划，01背包问题，这道题就是把01背包问题的物品价值从1维扩展到2维
*/
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < len(strs); i++ {
		zeroNum := 0
		oneNum := 0
		for _, c := range strs[i] {
			if c == '0' {
				zeroNum++
			} else {
				oneNum++
			}
		}
		for j := m; j >= zeroNum; j-- {
			for k := n; k >= oneNum; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zeroNum][k-oneNum]+1)
			}
		}
	}
	return dp[m][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
