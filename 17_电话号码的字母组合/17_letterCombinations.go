package _7_电话号码的字母组合

/*
1、使用回溯法
需要copy(temp, track)
2、横向for循环每次都从下标0开始，遍历每个数字代表的字母集合中的所有字母，纵向递归取不同的集合
*/

var letterMap = []string{
	"",
	"",
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}
var res []string

func letterCombinations(digits string) []string {
	res = []string{}
	if digits == "" {
		return res
	}
	track := make([]byte, 0, len(digits))
	backtrack(digits, track)
	return res
}

func backtrack(digits string, track []byte) {
	if len(track) == len(digits) {
		temp := make([]byte, len(track))
		copy(temp, track)
		res = append(res, string(temp))
		return
	}

	digit := digits[len(track)] - '0'
	letters := []byte(letterMap[digit])
	for i := 0; i < len(letters); i++ {
		track = append(track, letters[i])
		backtrack(digits, track)
		track = track[:len(track)-1]
	}
}
