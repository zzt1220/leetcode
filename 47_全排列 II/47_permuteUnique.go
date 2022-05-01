package _7_全排列_II

import "sort"

/*
1、使用used数组表示是否被使用过
2、全排列都是从0下标开始遍历
3、使用nums[i] == nums[i-1] && used[i-1] == false条件去重
*/

var res [][]int

func permuteUnique(nums []int) [][]int {
	res = [][]int{}
	used := make([]bool, len(nums))
	pathNums := make([]int, 0, len(nums))
	sort.Ints(nums)
	backtrack(nums, pathNums, used)
	return res
}

func backtrack(nums, pathNums []int, used []bool) {
	if len(nums) == len(pathNums) {
		tmp := make([]int, len(nums))
		copy(tmp, pathNums)
		res = append(res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
			continue
		}
		if !used[i] {
			used[i] = true
			pathNums = append(pathNums, nums[i])
			backtrack(nums, pathNums, used)
			pathNums = pathNums[:len(pathNums)-1]
			used[i] = false
		}
	}
}
