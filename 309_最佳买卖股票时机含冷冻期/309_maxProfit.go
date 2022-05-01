package _09_最佳买卖股票时机含冷冻期

/*
1、动态规划
2、dp[i][0]、dp[i][1]、dp[i][2]含义
dp[i][0]表示第i天买入股票状态的最大盈利
dp[i][1]表示第i天卖出股票状态（第i天当天卖出）的最大盈利
dp[i][2]表示第i天卖出股票状态（第i天之前就卖出）的最大盈利
3、滚动数组
*/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, 3)
	}

	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i%2][0] = max(dp[(i-1)%2][0], dp[(i-1)%2][2]-prices[i])
		dp[i%2][1] = dp[(i-1)%2][0] + prices[i]
		dp[i%2][2] = max(dp[(i-1)%2][2], dp[(i-1)%2][1])
	}
	return max(dp[(len(prices)-1)%2][1], dp[(len(prices)-1)%2][2])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
