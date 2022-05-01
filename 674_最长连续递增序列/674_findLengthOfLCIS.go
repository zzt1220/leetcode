package _74_最长连续递增序列

///*
//1、贪心算法
//*/
//func findLengthOfLCIS(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	count := 1
//	result := 1
//	for i := 1; i < len(nums); i++ {
//		if nums[i] > nums[i-1] {
//			count++
//		} else {
//			count = 1
//		}
//		if result < count {
//			result = count
//		}
//	}
//	return result
//}

/*
1、动态规划
*/

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = 1
	}
	result := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i%2] = dp[(i-1)%2] + 1
		} else {
			dp[i%2] = 1
		}
		if result < dp[i%2] {
			result = dp[i%2]
		}
	}
	return result
}
