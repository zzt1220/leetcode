package _0_子集_II

import "sort"

var res [][]int

func subsetsWithDup(nums []int) (ans [][]int) {
	res = [][]int{}
	var track []int
	used := make([]bool, len(nums))
	sort.Ints(nums)
	backtrack(nums, 0, used, track)
	return res
}
func backtrack(nums []int, start int, used []bool, track []int) {
	temp := make([]int, len(track))
	copy(temp, track)
	res = append(res, temp)
	for i := start; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
			continue
		}
		used[i] = true
		track = append(track, nums[i])
		backtrack(nums, i+1, used, track)
		track = track[:len(track)-1]
		used[i] = false
	}
}
