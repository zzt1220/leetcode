package _03_下一个更大元素_II

/*
1、单调栈
2、与第739道题类似
*/
func nextGreaterElements(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = -1
	}
	stack := []int{}
	for i := 0; i < len(nums)*2; i++ {
		for len(stack) > 0 && nums[i%len(nums)] > nums[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = nums[i%len(nums)]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%len(nums))
	}
	return ans
}
