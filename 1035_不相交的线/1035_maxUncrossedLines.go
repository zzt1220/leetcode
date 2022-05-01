package _035_不相交的线

/*
1、动态规划
2、本质上与第1143道题一样，就是求最长公共子序列的长度
*/
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, len(nums2)+1)
	}
	result := 0
	for i := 1; i < len(nums1)+1; i++ {
		for j := 1; j < len(nums2)+1; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i%2][j] = dp[(i-1)%2][j-1] + 1
				if result < dp[i%2][j] {
					result = dp[i%2][j]
				}
			} else {
				dp[i%2][j] = max(dp[(i-1)%2][j], dp[i%2][j-1])
			}
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
