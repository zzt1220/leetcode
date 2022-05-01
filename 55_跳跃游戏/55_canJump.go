package _5_跳跃游戏

/*
1、贪心算法，每次移动取最大跳跃步数（得到最大的覆盖范围），每移动一个单位，就更新最大覆盖范围
*/
func canJump(nums []int) bool {
	cover := 0
	for i := 0; i <= cover; i++ {
		cover = max(cover, i+nums[i])
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
