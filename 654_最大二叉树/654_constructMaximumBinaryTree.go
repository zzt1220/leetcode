package _54_最大二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
1、递归
*/
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxValIndex := findMaxValIndex(nums)
	node := new(TreeNode)
	node.Val = nums[maxValIndex]
	node.Left = constructMaximumBinaryTree(nums[:maxValIndex])
	node.Right = constructMaximumBinaryTree(nums[maxValIndex+1:])
	return node
}

func findMaxValIndex(nums []int) int {
	maxValIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[maxValIndex] {
			maxValIndex = i
		}
	}
	return maxValIndex
}
