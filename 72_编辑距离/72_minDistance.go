package _2_编辑距离

/*
1、动态规划，编辑距离
2、与第583道题区别：
第583道题只能删除操作，没有替换操作
本题有删除，增加（等同于删除），替换操作
所以两道题在实现上唯一区别在dp[i][j] = min(dp[i][j-1]+1, dp[i-1][j]+1, dp[i-1][j-1]+2)和dp[i][j]=min(dp[i][j-1],dp[i-1][j],dp[i-1][j-1])+1
*/
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i <= len(word1); i++ {
		dp[i][0] = i
	}
	for j := 0; j <= len(word2); j++ {
		dp[0][j] = j
	}
	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(x, y, z int) int {
	if x < y && x < z {
		return x
	}
	if y < z {
		return y
	}
	return z
}
