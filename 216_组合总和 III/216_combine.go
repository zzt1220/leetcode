package _16_组合总和_III

/*
1、使用回溯法
需要copy(temp, track)
2、使用剪枝优化，条件表达式：n-(k-len(track))+1和sum > targetSum
举个例子，n=4，k=3， 目前已经选取的元素为0，即len(track)为0，那么n-(k-0)+1就是4-(3-0)+1=2
*/

var res [][]int

func combinationSum3(k int, n int) [][]int {
	res = [][]int{}
	sum := 0
	track := make([]int, 0, k)
	backtrack(n, sum, k, 1, track)
	return res
}

func backtrack(targetSum, sum, k, start int, track []int) {
	if len(track) == k {
		if targetSum == sum {
			temp := make([]int, k)
			copy(temp, track)
			res = append(res, temp)
		}
		return
	}
	for i := start; i <= 9-(k-len(track))+1 && (sum+i <= targetSum); i++ {
		sum = sum + i
		track = append(track, i)
		backtrack(targetSum, sum, k, i+1, track)
		track = track[:len(track)-1]
		sum = sum - i
	}
}
