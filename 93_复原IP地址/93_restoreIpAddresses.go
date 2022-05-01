package _3_复原IP地址

/*
1、使用回溯法
*/

import "strconv"

var res []string

func restoreIpAddresses(s string) []string {
	res = []string{}
	var track []string
	backtrack(s, 0, track)
	return res
}
func backtrack(s string, start int, track []string) {
	if len(s) == start && len(track) == 4 {
		temp := track[0] + "." + track[1] + "." + track[2] + "." + track[3]
		res = append(res, temp)
		return
	}
	for i := start; i < len(s); i++ {
		if i-start < 3 && len(track) < 4 {
			if isNormalIp(s, start, i) {
				track = append(track, s[start:i+1])
			} else {
				continue
			}
		} else {
			break
		}
		backtrack(s, i+1, track)
		track = track[:len(track)-1]
	}
}
func isNormalIp(s string, start, end int) bool {
	checkInt, _ := strconv.Atoi(s[start : end+1])
	if end-start+1 > 1 && s[start] == '0' { //对于前导 0的IP（特别注意s[startIndex]=='0'的判断，不应该写成s[startIndex]==0，因为s截取出来不是数字）
		return false
	}
	if checkInt > 255 {
		return false
	}
	return true
}
