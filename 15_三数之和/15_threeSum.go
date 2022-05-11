package _5_三数之和

/*
1、双指针
*/
import "sort"

func threeSum(nums []int) [][]int {
	var ret [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		// 新增下面判断条件
		if nums[i] > 0 {
			break
		}
		// 不可以包含重复的三元组
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			// 不可以包含重复的三元组
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if nums[i]+nums[j]+nums[k] == 0 {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return ret
}
