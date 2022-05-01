package _14_买卖股票的最佳时机含手续费

/*
方法一：动态规划
1、与第122道题唯一区别就是卖出时候需要减去手续费
*/
func maxProfit(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		dp[i%2][0] = max(dp[(i-1)%2][0], dp[(i-1)%2][1]-prices[i])
		dp[i%2][1] = max(dp[(i-1)%2][1], dp[(i-1)%2][0]+prices[i]-fee)
	}
	return dp[(len(prices)-1)%2][1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
方法二：贪心算法
比较难理解
*/
