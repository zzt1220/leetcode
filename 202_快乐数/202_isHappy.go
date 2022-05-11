package _02_快乐数

/*
1、题目中说了可能会无限循环，但始终变不到1，也就是说求和的过程中，sum会重复出现
*/
func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 && !m[n] {
		n, m[n] = getSum(n), true
	}
	return n == 1
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
