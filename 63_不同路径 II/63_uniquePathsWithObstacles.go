package _3_不同路径_II

/*
1、动态规划
*/

//func uniquePathsWithObstacles(obstacleGrid [][]int) int {
//	m, n := len(obstacleGrid), len(obstacleGrid[0])
//	// 定义一个dp数组
//	dp := make([][]int, m)
//	for i := 0; i < m; i++ {
//		dp[i] = make([]int, n)
//	}
//	// 初始化
//	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
//		// 如果是障碍物, 后面的就都是0, 不用循环了
//		dp[i][0] = 1
//	}
//	for i := 0; i < n && obstacleGrid[0][i] == 0; i++ {
//		dp[0][i] = 1
//	}
//	// dp数组推导过程
//	for i := 1; i < m; i++ {
//		for j := 1; j < n; j++ {
//			if obstacleGrid[i][j] == 0 {
//				dp[i][j] = dp[i-1][j] + dp[i][j-1]
//			}
//		}
//	}
//	return dp[m-1][n-1]
//}

//滚动数组
//注意每行都需要对dp[0]做初始化操作
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, n)
	for i := 0; i < n && obstacleGrid[0][i] == 0; i++ {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		if dp[0] == 1 && obstacleGrid[i][0] == 0 {
			dp[0] = 1
		} else {
			dp[0] = 0
		}
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[j] = dp[j-1] + dp[j]
			} else {
				dp[j] = 0
			}
		}
	}
	return dp[n-1]
}
