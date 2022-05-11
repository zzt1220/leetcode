package _83_移动零

/*
1、双指针
2、注意：slowIndex != fastIndex这边条件下，才能对fastIndex下标元素赋0值
*/
func moveZeroes(nums []int) {
	slowIndex := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != 0 {
			nums[slowIndex] = nums[fastIndex]
			if slowIndex != fastIndex {
				nums[fastIndex] = 0
			}
			slowIndex++
		}
	}
}
