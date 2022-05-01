package _31_分割回文串

/*
1、使用回溯法
需要copy(temp, track)
2、使用剪枝优化，条件表达式：n-(k-len(track))+1
举个例子，n=4，k=3， 目前已经选取的元素为0，即len(track)为0，那么n-(k-0)+1就是4-(3-0)+1=2
*/

var res [][]string

func partition(s string) [][]string {
	res = [][]string{}
	var track []string
	backtrack(s, 0, track)
	return res
}
func backtrack(s string, start int, track []string) {
	if len(s) == start {
		temp := make([]string, len(track))
		copy(temp, track)
		res = append(res, temp)
		return
	}
	for i := start; i < len(s); i++ {
		if isPalindrome(s, start, i) {
			track = append(track, s[start:i+1])
		} else {
			continue
		}
		backtrack(s, i+1, track)
		track = track[:len(track)-1]
	}
}
func isPalindrome(s string, start, end int) bool {
	left := start
	right := end
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
