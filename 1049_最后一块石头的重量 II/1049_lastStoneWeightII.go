package _049_最后一块石头的重量_II

/*
1、动态规划。01背包，与第416道题思想一样
2、滚动数组
*/
func lastStoneWeightII(stones []int) int {
	sum := 0
	for i := 0; i < len(stones); i++ {
		sum += stones[i]
	}
	target := sum / 2

	dp := make([]int, target+1)
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return sum - dp[target] - dp[target]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
