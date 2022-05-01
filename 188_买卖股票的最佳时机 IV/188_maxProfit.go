package _88_买卖股票的最佳时机_IV

/*
1、动态规划
2、第123道题的扩展
*/

func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	if k == 0 {
		return 0
	}

	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, 2*k)
	}

	for i := 0; i < k; i++ {
		dp[0][i*2] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		dp[i%2][0] = max(dp[(i-1)%2][0], -prices[i])
		dp[i%2][1] = max(dp[(i-1)%2][1], dp[(i-1)%2][0]+prices[i])
		for j := 1; j < k; j++ {
			dp[i%2][j*2] = max(dp[(i-1)%2][j*2], dp[(i-1)%2][j*2-1]-prices[i])
			dp[i%2][j*2+1] = max(dp[(i-1)%2][j*2+1], dp[(i-1)%2][j*2]+prices[i])
		}
	}
	return dp[(len(prices)-1)%2][k*2-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
