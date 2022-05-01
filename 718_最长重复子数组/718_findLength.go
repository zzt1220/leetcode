package _18_最长重复子数组

///*
//1、动态规划（不使用滚动数组）
//*/
//func findLength(nums1 []int, nums2 []int) int {
//	if len(nums1) == 0 || len(nums2) == 0 {
//		return 0
//	}
//	dp := make([][]int, len(nums1)+1)
//	for i := 0; i < len(nums1)+1; i++ {
//		dp[i] = make([]int, len(nums2)+1)
//	}
//
//	result := 0
//	for i := 1; i <= len(nums1); i++ {
//		for j := 1; j <= len(nums2); j++ {
//			if nums1[i-1] == nums2[j-1] {
//				dp[i][j] = dp[i-1][j-1] + 1
//				if result < dp[i][j] {
//					result = dp[i][j]
//				}
//			}
//		}
//	}
//	return result
//}

/*
1、动态规划（使用滚动数组）
*/
func findLength(nums1 []int, nums2 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}
	dp := make([]int, len(nums2)+1)

	result := 0
	for i := 1; i <= len(nums1); i++ {
		for j := len(nums2); j >= 1; j-- {
			if nums1[i-1] == nums2[j-1] {
				dp[j] = dp[j-1] + 1
				if result < dp[j] {
					result = dp[j]
				}
			} else {
				dp[j] = 0
			}
		}
	}
	return result
}
