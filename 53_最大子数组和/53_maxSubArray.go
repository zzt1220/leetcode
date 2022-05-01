package _3_最大子数组和

///*
//1、贪心算法
//*/
//
//func maxSubArray(nums []int) int {
//	count := 0
//	result := nums[0]
//	for i := 0; i < len(nums); i++ {
//		count += nums[i]
//		if result < count {
//			result = count
//		}
//		if count < 0 {
//			count = 0
//		}
//	}
//	return result
//}

/*
2、动态规划
*/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, 2)
	dp[0] = nums[0]
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i%2] = max(nums[i], dp[(i-1)%2]+nums[i])
		if result < dp[i%2] {
			result = dp[i%2]
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
