package _46_使用最小花费爬楼梯

/*
1、动态规划
*/
func minCostClimbingStairs(cost []int) int {
	if len(cost) < 2 {
		return 0
	}

	var res int
	var dp [2]int
	dp[0] = 0
	dp[1] = 0
	for i := 2; i < len(cost)+1; i++ {
		res = min(dp[1]+cost[i-1], dp[0]+cost[i-2])
		dp[0] = dp[1]
		dp[1] = res
	}
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
