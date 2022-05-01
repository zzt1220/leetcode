package _22_买卖股票的最佳时机_II

////方法一：贪心算法
///*
//1、贪心算法，正利润加起来
//*/
//
//func maxProfit(prices []int) int {
//	profit := 0
//	for i := 1; i < len(prices); i++ {
//		if prices[i] > prices[i-1] {
//			profit = profit + prices[i] - prices[i-1]
//		}
//	}
//	return profit
//}

//方法二：动态规划
/*
1、dp[i][0]和dp[i][1]的含义
dp[i][0]表示第i天持有股票的最大现金，注意与第121道题的区别
dp[i][1]表示第i天不持有股票的最大现金
2、使用滚动数组
*/
func maxProfit(prices []int) int {
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
		dp[i%2][1] = max(dp[(i-1)%2][1], dp[(i-1)%2][0]+prices[i])
	}
	return dp[(len(prices)-1)%2][1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
