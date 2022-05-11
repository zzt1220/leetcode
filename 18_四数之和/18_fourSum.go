package _8_四数之和

/*
1、与第15道题类似解法
*/

import "sort"

func fourSum(nums []int, target int) [][]int {
	var ret [][]int
	sort.Ints(nums)
	for m := 0; m < len(nums)-3; m++ {
		// 不能使用下面判断
		// if nums[m]>target{
		//     break
		// }
		if m > 0 && nums[m] == nums[m-1] {
			continue
		}
		for i := m + 1; i < len(nums)-2; i++ {
			if i > m+1 && nums[i] == nums[i-1] {
				continue
			}
			j, k := i+1, len(nums)-1
			for j < k {
				if j > i+1 && nums[j] == nums[j-1] {
					j++
					continue
				}
				if nums[m]+nums[i]+nums[j]+nums[k] == target {
					ret = append(ret, []int{nums[m], nums[i], nums[j], nums[k]})
					j++
					k--
				} else if nums[m]+nums[i]+nums[j]+nums[k] < target {
					j++
				} else {
					k--
				}
			}
		}
	}
	return ret
}
