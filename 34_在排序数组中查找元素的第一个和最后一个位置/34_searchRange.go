package _4_在排序数组中查找元素的第一个和最后一个位置

func searchRange(nums []int, target int) []int {
	pos := binarySearch(nums, target)
	if pos == -1 {
		return []int{-1, -1}
	}
	start := pos
	end := pos
	for start > 0 {
		if nums[start] != nums[start-1] {
			break
		}
		start--
	}
	for end < len(nums)-1 {
		if nums[end] != nums[end+1] {
			break
		}
		end++
	}
	return []int{start, end}
}

func binarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
