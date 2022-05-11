package _4_螺旋矩阵

/*
1、方法一：使用j <= maxX和j <= maxY作为上边for循环和右边for循环的结束条件
2、下边for循环和左边for循环执行前需要判断maxY-minY >= 1和maxX-minX >= 1
3、可以与第59道题一起看
*/
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	l1 := (m + 1) / 2
	l2 := (n + 1) / 2
	ret := make([]int, m*n)
	var index int
	var round int
	if l1 > l2 {
		round = l2
	} else {
		round = l1
	}
	for i := 0; i < round; i++ {
		minX, minY, maxX, maxY := i, i, n-i-1, m-i-1
		for j := minX; j <= maxX; j++ {
			ret[index] = matrix[minY][j]
			index++
		}
		for j := minY + 1; j <= maxY; j++ {
			ret[index] = matrix[j][maxX]
			index++
		}
		if maxY-minY >= 1 {
			for j := maxX - 1; j > minX; j-- {
				ret[index] = matrix[maxY][j]
				index++
			}
		}
		if maxX-minX >= 1 {
			for j := maxY; j > minY; j-- {
				ret[index] = matrix[j][minX]
				index++
			}
		}

	}
	return ret
}

///*
//1、方法二：使用j < maxX和j < maxY作为上边for循环和右边for循环的结束条件
//2、与方法一区别就是最后需要判断是否需要对最后一个值进行赋值，所以麻烦一些，建议直接使用方法一
//*/
//func spiralOrder(matrix [][]int) []int {
//	m := len(matrix)
//	if m == 0 {
//		return nil
//	}
//	n := len(matrix[0])
//	l1 := (m + 1) / 2
//	l2 := (n + 1) / 2
//	ret := make([]int, m*n)
//	var index int
//	var round int
//	if l1 > l2 {
//		round = l2
//	} else {
//		round = l1
//	}
//
//	for i := 0; i < round; i++ {
//		minX, minY, maxX, maxY := i, i, n-i-1, m-i-1
//		for j := minX; j < maxX; j++ {
//			ret[index] = matrix[minY][j]
//			index++
//		}
//		for j := minY; j < maxY; j++ {
//			ret[index] = matrix[j][maxX]
//			index++
//		}
//		if maxY-minY >= 1 {
//			for j := maxX; j > minX; j-- {
//				ret[index] = matrix[maxY][j]
//				index++
//			}
//		}
//		if maxX-minX >= 1 {
//			for j := maxY; j > minY; j-- {
//				ret[index] = matrix[j][minX]
//				index++
//			}
//		}
//	}
//	if m <= n && m%2 == 1 {
//		ret[m*n-1] = matrix[m/2][n-1-m/2]
//	}
//	if m > n && n%2 == 1 {
//		ret[m*n-1] = matrix[m-1-n/2][n/2]
//	}
//	return ret
//}
