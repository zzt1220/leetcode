package _55_分发饼干

/*
1、贪心算法，一种方法从胃口大开始分配，这时候遍历人，一种是从胃口小开始分配，这时候开始遍历饼干
*/

import "sort"

//func findContentChildren(g []int, s []int) int {
//	sort.Ints(g)
//	sort.Ints(s)
//	index := len(s) - 1
//	result := 0
//	for i := len(g) - 1; i >= 0 && index >= 0; i-- {
//		if g[i] <= s[index] {
//			result++
//			index--
//		}
//	}
//	return result
//}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	index := 0
	result := 0
	for i := 0; i < len(s) && index < len(g); i++ {
		if g[index] <= s[i] {
			result++
			index++
		}
	}
	return result
}
