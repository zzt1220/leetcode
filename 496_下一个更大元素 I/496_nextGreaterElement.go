package _96_下一个更大元素_I

/*
1、单调栈
2、与第739道题类似
*/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		ans[i] = -1
	}
	num1map := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		num1map[nums1[i]] = i
	}
	stack := []int{}
	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			if index, exist := num1map[nums2[stack[len(stack)-1]]]; exist {
				ans[index] = nums2[i]
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
