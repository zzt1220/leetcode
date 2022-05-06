package _10_平衡二叉树

/*
1、方法一：递归
2、按照后序遍历的思想
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if nodeHeight(root) == -1 {
		return false
	}
	return true
}

func nodeHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftH := nodeHeight(node.Left)
	if leftH == -1 {
		return -1
	}
	rightH := nodeHeight(node.Right)
	if rightH == -1 {
		return -1
	}
	if abs(nodeHeight(node.Left), nodeHeight(node.Right)) > 1 {
		return -1
	}
	return max(nodeHeight(node.Left), nodeHeight(node.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
