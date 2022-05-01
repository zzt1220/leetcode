package _0_组合总和_II

/*
1、使用回溯法
需要copy(temp, track)
2、使用剪枝优化，条件表达式：(sum+candidates[i]<=targetSum)
注意：需要先对数集合进行从小到大的排序，然后才能做剪枝操作
3、used[i-1] == false，说明同一树层candidates[i-1]使用过，题目要求不能有重复组合，所以直接continue
*/

import "sort"

var res [][]int

func combinationSum2(candidates []int, target int) [][]int {
	res = [][]int{}
	used := make([]bool, len(candidates))
	track := []int{}
	sort.Ints(candidates)
	backtrack(candidates, target, 0, 0, track, used)
	return res
}

func backtrack(candidates []int, targetSum, sum, start int, track []int, used []bool) {
	if targetSum == sum {
		temp := make([]int, len(track))
		copy(temp, track)
		res = append(res, temp)
		return
	}

	for i := start; i < len(candidates) && (sum+candidates[i] <= targetSum); i++ {
		if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
			continue
		}
		sum = sum + candidates[i]
		track = append(track, candidates[i])
		used[i] = true
		backtrack(candidates, targetSum, sum, i+1, track, used)
		used[i] = false
		track = track[:len(track)-1]
		sum = sum - candidates[i]
	}
}
