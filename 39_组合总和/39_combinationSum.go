package _9_组合总和

/*
1、使用回溯法
需要copy(temp, track)
2、使用剪枝优化，条件表达式：(sum+candidates[i]<=targetSum)
注意：需要先对数集合进行从小到大的排序，然后才能做剪枝操作
*/

import "sort"

var res [][]int

func combinationSum(candidates []int, target int) [][]int {
	res = [][]int{}
	sort.Ints(candidates)
	track := make([]int, 0, len(candidates))
	backtrack(candidates, target, 0, 0, track)
	return res
}

func backtrack(candidates []int, targetSum, sum, start int, track []int) {
	if targetSum == sum {
		temp := make([]int, len(track))
		copy(temp, track)
		res = append(res, temp)
		return
	}

	for i := start; i < len(candidates) && (sum+candidates[i] <= targetSum); i++ {
		sum = sum + candidates[i]
		track = append(track, candidates[i])
		backtrack(candidates, targetSum, sum, i, track)
		track = track[:len(track)-1]
		sum = sum - candidates[i]
	}
}
