package _16_最长回文子序列

///*
//1、动态规划（没有使用滚动数组优化）
//2、注意本题是最长回文子序列，与第647题回文子串不一样
//*/
//func longestPalindromeSubseq(s string) int {
//	dp := make([][]int, len(s))
//	for i := 0; i < len(s); i++ {
//		dp[i] = make([]int, len(s))
//	}
//	for i := 0; i < len(s); i++ {
//		dp[i][i] = 1
//	}
//	for i := len(s) - 1; i >= 0; i-- {
//		for j := i + 1; j < len(s); j++ {
//			if s[i] == s[j] {
//				dp[i][j] = dp[i+1][j-1] + 2
//			} else {
//				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
//			}
//		}
//	}
//	return dp[0][len(s)-1]
//}
//
//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}

/*
1、动态规划（使用滚动数组优化）
2、因为使用滚动数组，所以每层循环需要初始化dp[i%2][i] = 1
3、这里滚动数组不好理解，可以使用上面非滚动数组解法
*/

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		dp[i%2][i] = 1
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i%2][j] = dp[(i+1)%2][j-1] + 2
			} else {
				dp[i%2][j] = max(dp[(i+1)%2][j], dp[i%2][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
