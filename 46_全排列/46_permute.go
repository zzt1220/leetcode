package _6_全排列

/*
1、使用used数组表示是否被使用过
2、全排列都是从0下标开始遍历
*/

var res [][]int

func permute(nums []int) [][]int {
	res = [][]int{}
	used := make([]bool, len(nums))
	pathNums := make([]int, 0, len(nums))
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
		if !used[i] {
			used[i] = true
			pathNums = append(pathNums, nums[i])
			backtrack(nums, pathNums, used)
			pathNums = pathNums[:len(pathNums)-1]
			used[i] = false
		}
	}
}
