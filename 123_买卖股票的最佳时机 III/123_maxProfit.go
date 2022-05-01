package _23_买卖股票的最佳时机_III

/*
1、动态规划
2、dp数组含义：
dp[i][1]表示第i天第1次买入
dp[i][2]表示第i天第1次卖出
dp[i][3]表示第i天第2次买入
dp[i][3]表示第i天第2次卖出
3、第121,122,123道题区别：
第121道题是不限次数买卖
第122道题是限制只能买卖一次
第123道题是限制只能买卖两次
*/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, 4)
	}

	dp[0][0] = -prices[0]
	dp[0][2] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i%2][0] = max(dp[(i-1)%2][0], -prices[i])
		dp[i%2][1] = max(dp[(i-1)%2][1], dp[(i-1)%2][0]+prices[i])
		dp[i%2][2] = max(dp[(i-1)%2][2], dp[(i-1)%2][1]-prices[i])
		dp[i%2][3] = max(dp[(i-1)%2][3], dp[(i-1)%2][2]+prices[i])
	}
	return dp[(len(prices)-1)%2][3]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
