package _63_划分字母区间

/*
1、统计每一个字符最后出现的位置
2、从头遍历字符，如果找到之前遍历过的所有字母的最远边界，说明这个边界就是分割点
*/
func partitionLabels(s string) []int {
	var position [26]int
	for i := 0; i < len(s); i++ {
		position[s[i]-'a'] = i
	}
	left := 0
	right := 0
	var result []int
	for i := 0; i < len(s); i++ {
		right = max(right, position[s[i]-'a'])
		if right == i {
			result = append(result, right-left+1)
			left = i + 1
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
