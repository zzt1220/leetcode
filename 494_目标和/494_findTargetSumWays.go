package _94_目标和

/*
1、动态规划，01背包问题
2、推导
left - right = target
left + right = sum
left = (target + sum)/2
问题转化为在集合nums中找出和为left的组合个数
*/
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if target > sum || target < (-sum) {
		return 0
	}
	if (target+sum)%2 == 1 {
		return 0
	}

	bagSize := (target + sum) / 2
	dp := make([]int, bagSize+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := bagSize; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[bagSize]
}
