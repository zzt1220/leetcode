package _83_两个字符串的删除操作

///*
//方法一：通过先求最长公共子序列来实现，见第1143道题
//1、动态规划
//*/
//func minDistance(word1 string, word2 string) int {
//	dp := make([][]int, len(word1)+1)
//	for i := 0; i < len(word1)+1; i++ {
//		dp[i] = make([]int, len(word2)+1)
//	}
//	result := 0
//	for i := 1; i < len(word1)+1; i++ {
//		for j := 1; j < len(word2)+1; j++ {
//			if word1[i-1] == word2[j-1] {
//				dp[i][j] = dp[i-1][j-1] + 1
//				if result < dp[i][j] {
//					result = dp[i][j]
//				}
//			} else {
//				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
//			}
//		}
//	}
//	return len(word1) + len(word2) - result*2
//}
//
//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}

/*
方法二：直接使用动态规划
1、动态规划
2、dp[i][j]含义：
以i-1下标结尾的字符串word1，和以j-1下标结尾的字符串word2，要使得两个字符串相同所需的最小步数
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
				dp[i][j] = min(dp[i][j-1]+1, dp[i-1][j]+1, dp[i-1][j-1]+2)
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
