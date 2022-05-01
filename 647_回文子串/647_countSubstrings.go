package _47_回文子串

///*
//1、方法一：双指针法
//2、复杂度
//时间复杂度：O(n^2)
//空间复杂度：O(1)
//*/
//func countSubstrings(s string) int {
//	result := 0
//	for i := 0; i < len(s); i++ {
//		result += extend(s, i, i)
//		result += extend(s, i, i+1)
//	}
//	return result
//}
//
//func extend(s string, i int, j int) int {
//	res := 0
//	for i >= 0 && j < len(s) && (s[i] == s[j]) {
//		res++
//		i--
//		j++
//	}
//	return res
//}

/*
1、方法二：动态规划
2、注意这里dp数组的含义是是否是回文子串，而不是回文子串的数目
3、复杂度
时间复杂度：O(n^2)
空间复杂度：O(n^2)
*/
func countSubstrings(s string) int {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					result++
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					result++
					dp[i][j] = true
				}
			}
		}
	}
	return result
}
