package 剑指_Offer_58___II__左旋转字符串

/*
1、反转前n个字符
2、反转第n到end字符
3、反转整个字符
4、与第151道题类似
*/
func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	reverse(b, 0, n-1)
	reverse(b, n, len(b)-1)
	reverse(b, 0, len(b)-1)
	return string(b)
}

func reverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}
