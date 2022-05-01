package _5_跳跃游戏_II

/*
1、贪心算法
*/
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	curCover := 0
	nextCover := 0
	step := 0
	for i := 0; i < len(nums); i++ {
		if nextCover < nums[i]+i {
			nextCover = nums[i] + i
			if nextCover >= len(nums)-1 {
				step++
				break
			}
		}
		if i == curCover && i != len(nums)-1 {
			step++
			curCover = nextCover
		}
	}
	return step
}
