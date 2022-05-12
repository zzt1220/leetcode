package _2_接雨水

/*
1、方法一：双指针
2、按列计算
*/
func trap(height []int) int {
	sum := 0
	for i := 1; i < len(height)-1; i++ {
		lHeight := height[i]
		rHeight := height[i]
		for j := i - 1; j >= 0; j-- {
			if lHeight < height[j] {
				lHeight = height[j]
			}
		}
		for j := i + 1; j < len(height); j++ {
			if rHeight < height[j] {
				rHeight = height[j]
			}
		}
		h := min(lHeight, rHeight) - height[i]
		sum = sum + h
	}
	return sum
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

/*
1、方法二：动态规划计算左边最大值和右边最大值
2、按列计算
*/
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	sum := 0
	maxLeft := make([]int, len(height))
	maxRight := make([]int, len(height))
	maxLeft[0] = height[0]
	maxRight[len(height)-1] = height[len(height)-1]
	for i := 1; i < len(height); i++ {
		maxLeft[i] = max(height[i], maxLeft[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		maxRight[i] = max(height[i], maxRight[i+1])
	}
	for i := 1; i < len(height)-1; i++ {
		h := min(maxLeft[i], maxRight[i]) - height[i]
		sum = sum + h
	}
	return sum
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

/*
1、方法三：单调栈（栈底到栈顶从大到小）
2、按行计算
*/
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	sum := 0
	stack := []int{}
	stack = append(stack, 0)
	for i := 1; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				h := min(height[stack[len(stack)-1]], height[i]) - height[mid]
				w := i - stack[len(stack)-1] - 1
				sum = sum + w*h
			}
		}
		stack = append(stack, i)
	}
	return sum
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
