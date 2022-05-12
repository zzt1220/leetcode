package _4_柱状图中最大的矩形

/*
1、方法一：双指针(执行会超时)
*/
func largestRectangleArea(heights []int) int {
	sum := 0
	for i := 0; i < len(heights); i++ {
		left := i
		right := i
		for left >= 0 {
			if heights[left] < heights[i] {
				break
			}
			left--
		}
		for right < len(heights) {
			if heights[right] < heights[i] {
				break
			}
			right++
		}
		w := right - left - 1
		sum = max(sum, w*heights[i])
	}
	return sum
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
1、方法二：动态规划
2、注意for t>=0&&heights[t]>=heights[i]这里是for循环，不是if
*/
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	sum := 0
	minLeftIndex := make([]int, len(heights))
	minRightIndex := make([]int, len(heights))
	minLeftIndex[0] = -1
	for i := 1; i < len(heights); i++ {
		t := i - 1
		for t >= 0 && heights[t] >= heights[i] {
			t = minLeftIndex[t]
		}
		minLeftIndex[i] = t
	}
	minRightIndex[len(heights)-1] = len(heights)
	for i := len(heights) - 2; i >= 0; i-- {
		t := i + 1
		for t < len(heights) && heights[t] >= heights[i] {
			t = minRightIndex[t]
		}
		minRightIndex[i] = t
	}
	for i := 0; i < len(heights); i++ {
		w := minRightIndex[i] - minLeftIndex[i] - 1
		sum = max(sum, w*heights[i])
	}
	return sum
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
1、方法三：单调栈（栈底到栈顶从小到大）
2、在首部和尾部加上0值元素
3、与第42道题类似
*/
func largestRectangleArea(heights []int) int {
	stack := []int{}
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	stack = append(stack, 0)
	sum := 0
	for i := 1; i < len(heights); i++ {
		for heights[i] < heights[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			w := i - stack[len(stack)-1] - 1
			sum = max(sum, w*heights[mid])
		}
		stack = append(stack, i)
	}
	return sum
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
