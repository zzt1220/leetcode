package _9_螺旋矩阵_II

/*
1、方法一：使用j <= maxX和j <= maxY作为上边for循环和右边for循环的结束条件
*/
func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	round := (n + 1) / 2
	index := 1
	for i := 0; i < round; i++ {
		minX, minY, maxX, maxY := i, i, n-i-1, n-i-1
		for j := minX; j <= maxX; j++ {
			ret[minY][j] = index
			index++
		}
		for j := minY + 1; j <= maxY; j++ {
			ret[j][maxX] = index
			index++
		}
		for j := maxX - 1; j > minX; j-- {
			ret[maxY][j] = index
			index++
		}
		for j := maxY; j > minY; j-- {
			ret[j][minX] = index
			index++
		}
	}
	return ret
}

///*
//1、方法二：使用j < maxX和j < maxY作为上边for循环和右边for循环的结束条件
//2、与方法一区别就是最后需要判断是否需要对最后一个值进行赋值，所以麻烦一些，建议直接使用方法一
//*/
//func generateMatrix(n int) [][]int {
//	ret := make([][]int, n)
//	for i := 0; i < n; i++ {
//		ret[i] = make([]int, n)
//	}
//	round := (n + 1) / 2
//	index := 1
//	for i := 0; i < round; i++ {
//		minX, minY, maxX, maxY := i, i, n-i-1, n-i-1
//		for j := minX; j < maxX; j++ {
//			ret[minY][j] = index
//			index++
//		}
//		for j := minY; j < maxY; j++ {
//			ret[j][maxX] = index
//			index++
//		}
//		for j := maxX; j > minX; j-- {
//			ret[maxY][j] = index
//			index++
//		}
//		for j := maxY; j > minY; j-- {
//			ret[j][minX] = index
//			index++
//		}
//	}
//	if n%2 == 1 {
//		ret[n/2][n/2] = index
//	}
//	return ret
//}
