package _60_柠檬水找零

/*
1、贪心算法
*/

func lemonadeChange(bills []int) bool {
	five := 0
	ten := 0
	for _, bill := range bills {
		if bill == 5 {
			five++
		}
		if bill == 10 {
			if five > 0 {
				five--
				ten++
			} else {
				return false
			}
		}
		if bill == 20 {
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five > 2 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
