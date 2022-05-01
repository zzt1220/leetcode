package _7_组合

/*
1、使用回溯法
需要copy(temp, track)
2、使用剪枝优化，条件表达式：n-(k-len(track))+1
举个例子，n=4，k=3， 目前已经选取的元素为0，即len(track)为0，那么n-(k-0)+1就是4-(3-0)+1=2
*/

var res [][]int

func combine(n int, k int) [][]int {
	res = [][]int{}
	if n <= 0 || k <= 0 || k > n {
		return res
	}
	track := make([]int, 0, k)
	backtrack(n, k, 1, track)
	return res
}
func backtrack(n, k, start int, track []int) {
	if len(track) == k {
		temp := make([]int, k)
		copy(temp, track)
		res = append(res, temp)
		return
	}
	for i := start; i <= n-(k-len(track))+1; i++ {
		track = append(track, i)
		backtrack(n, k, i+1, track)
		track = track[:len(track)-1]
	}
}
