package _04_水果成篮

/*
1、滑动窗口
2、与第209道题相似
3、go的map类型有len方法可以知道目前有多少个kv
*/
func totalFruit(fruits []int) int {
	i, width, win := 0, 0, make(map[int]int) // i 是慢指针位置, width 是窗口最大宽度, win是用map来存放窗口内的种类
	for j := 0; j < len(fruits); j++ {       //对数组做一次遍历
		win[fruits[j]]++   // 每次都将快指针对应的元素添加进当前窗口内,如果有的话,更新其值,
		for len(win) > 2 { // 如果当前窗口元素种类大于2
			win[fruits[i]]--         // 从窗口里将该元素的种类对应的个数减1
			if win[fruits[i]] == 0 { // 如果对应个数为 0
				delete(win, fruits[i]) // 则把该元素删除
			}
			i++ // 慢指针往右移，目的是使窗口内的元素种类不超过2种
		} // 上面的循环结束后，确保当前窗口内的元素种类不超过2种
		tmpw := j - i + 1 // 查看当前窗口宽度大小
		if tmpw > width { // 如果当前窗口比之前的宽度大
			width = tmpw // 更新窗口
		}
	}
	return width // 返回更新过的窗口宽度
}
