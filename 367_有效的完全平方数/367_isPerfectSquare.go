package _67_有效的完全平方数

func isPerfectSquare(num int) bool {
	low := 0
	high := num
	for low <= high {
		mid := (low + high) / 2
		tmp := mid * mid
		if tmp < num {
			low = mid + 1
		} else if tmp > num {
			high = mid - 1
		} else {
			return true
		}
	}
	return false
}
