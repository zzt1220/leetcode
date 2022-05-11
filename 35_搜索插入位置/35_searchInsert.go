package _5_搜索插入位置

/*
1、for循环使用条件low < high，low == high情况留在循环外处理
*/
func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := (low + high) / 2
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	if nums[low] == target {
		return low
	}
	if nums[low] < target {
		return low + 1
	}
	return low
}
