package _38_单调递增的数字

/*
1、贪心算法
2、注意需要把字符串nstr转成byte切片，因为字符串不能修改，切片才可以修改
3、使用strconv要注意strconv.Itoa和strconv.Atoi返回参数个数不一样
*/

import "strconv"

func monotoneIncreasingDigits(n int) int {
	nstr := strconv.Itoa(n)
	nbyte := []byte(nstr)
	flag := len(nbyte)
	for i := len(nbyte) - 1; i > 0; i-- {
		if nbyte[i-1] > nbyte[i] {
			nbyte[i-1]--
			flag = i
		}
	}
	for i := flag; i < len(nbyte); i++ {
		nbyte[i] = '9'
	}
	result, _ := strconv.Atoi(string(nbyte))
	return result
}
