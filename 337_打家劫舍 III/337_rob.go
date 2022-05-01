package _37_打家劫舍_III

/*
1、动态规划+二叉树后序遍历
2、递归返回两个元素，一个是偷当前节点，一个是不偷当前节点
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	result := robTree(root)
	return max(result[0], result[1])
}

func robTree(cur *TreeNode) []int {
	if cur == nil {
		return []int{0, 0}
	}
	left := robTree(cur.Left)
	right := robTree(cur.Right)

	val1 := max(left[0], left[1]) + max(right[0], right[1])
	val2 := cur.Val + left[0] + right[0]

	return []int{val1, val2}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
