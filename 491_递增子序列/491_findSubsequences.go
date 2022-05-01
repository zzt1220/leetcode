package _91_递增子序列

/*
1、递增子序列的长度必须大于等于2
2、需要去重，注意不能排序后再去重，所以跟组合或子集去重处理不一样，这里使用map或组数来标记同一层中的元素是否出现过，如果出现过，说明被使用过，直接跳过
*/

var res [][]int

func findSubsequences(nums []int) [][]int {
	res = [][]int{}
	var track []int
	backtrack(nums, 0, track)
	return res
}
func backtrack(nums []int, start int, track []int) {
	if len(track) >= 2 {
		temp := make([]int, len(track))
		copy(temp, track)
		res = append(res, temp)
	}
	//used := make(map[int]bool)
	used := [201]int{}
	for i := start; i < len(nums); i++ {
		//if len(track) > 0 && nums[i] < track[len(track)-1] || used[nums[i]] == true {
		//	continue
		//}
		if len(track) > 0 && nums[i] < track[len(track)-1] || used[nums[i]+100] == 1 {
			continue
		}
		//used[nums[i]] = true
		used[nums[i]+100] = 1
		track = append(track, nums[i])
		backtrack(nums, i+1, track)
		track = track[:len(track)-1]
	}
}
