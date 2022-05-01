package _39_单词拆分

/*
1、动态规划，完全背包问题
2、两种遍历顺序都可以，这里最好使用外层循环遍历背包容量，内层循环遍历物品方法
*/
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		wordDictSet[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
