package _6_删除有序数组中的重复项

/*
1、双指针
*/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slowIndex := 1
	for fastIndex := 1; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != nums[fastIndex-1] {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}
	return slowIndex
}
