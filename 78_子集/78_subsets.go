package _8_子集

/*
1、使用回溯法
需要copy(temp, track)
2、使用剪枝优化，条件表达式：n-(k-len(track))+1
举个例子，n=4，k=3， 目前已经选取的元素为0，即len(track)为0，那么n-(k-0)+1就是4-(3-0)+1=2
*/

var res [][]int

func subsets(nums []int) (ans [][]int) {
	res = [][]int{}
	var track []int
	backtrack(nums, 0, track)
	return res
}
func backtrack(nums []int, start int, track []int) {
	temp := make([]int, len(track))
	copy(temp, track)
	res = append(res, temp)
	for i := start; i < len(nums); i++ {
		track = append(track, nums[i])
		backtrack(nums, i+1, track)
		track = track[:len(track)-1]
	}
}
