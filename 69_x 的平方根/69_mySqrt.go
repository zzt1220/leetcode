package _9_x_的平方根

/*
1、因为题中说：“返回类型是整数，结果只保留整数部分，小数部分将被舍去”，所以结果的平方一定小于或等于x
*/
func mySqrt(x int) int {
	low := 0
	high := x
	res := 0
	for low <= high {
		mid := (low + high) / 2
		if mid*mid <= x {
			res = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return res
}
