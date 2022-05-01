package _143_最长公共子序列

/*
1、动态规划
2、与第718道题的区别：
第718道题要求连续，本题不要求连续，所以当text1[i-1]！=text2[j-1]时，本题dp[i][j]=max(dp[i][j-1],dp[i-1][j])
*/
func longestCommonSubsequence(text1 string, text2 string) int {
	if text1 == "" || text2 == "" {
		return 0
	}

	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(text1)+1; i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	result := 0
	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if result < dp[i][j] {
					result = dp[i][j]
				}
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
