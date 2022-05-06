package _12_路径总和

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

///*
//1、方法一：递归
//*/
//func hasPathSum(root *TreeNode, targetSum int) bool {
//	if root == nil {
//		return false
//	}
//	return traversal(root, targetSum, 0)
//}
//
//func traversal(root *TreeNode, targetSum int, curSum int) bool {
//	curSum += root.Val
//	if root.Left == nil && root.Right == nil {
//		if curSum == targetSum {
//			return true
//		}
//	}
//	if root.Left != nil {
//		left := traversal(root.Left, targetSum, curSum)
//		if left {
//			return true
//		}
//	}
//	if root.Right != nil {
//		right := traversal(root.Right, targetSum, curSum)
//		if right {
//			return true
//		}
//	}
//	return false
//}

/*
1、方法二：迭代
*/
type NodePathSum struct {
	Node *TreeNode
	Sum  int
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	stack := list.New()
	rootSum := new(NodePathSum)
	rootSum.Node = root
	rootSum.Sum = root.Val
	stack.PushBack(rootSum)
	for stack.Len() > 0 {
		e := stack.Back()
		nodePathSum := e.Value.(*NodePathSum)
		stack.Remove(e)
		if nodePathSum.Node.Left == nil && nodePathSum.Node.Right == nil && nodePathSum.Sum == targetSum {
			return true
		}
		if nodePathSum.Node.Left != nil {
			nodePathSumLeft := new(NodePathSum)
			nodePathSumLeft.Node = nodePathSum.Node.Left
			nodePathSumLeft.Sum = nodePathSum.Sum + nodePathSum.Node.Left.Val
			stack.PushBack(nodePathSumLeft)
		}
		if nodePathSum.Node.Right != nil {
			nodePathSumRight := new(NodePathSum)
			nodePathSumRight.Node = nodePathSum.Node.Right
			nodePathSumRight.Sum = nodePathSum.Sum + nodePathSum.Node.Right.Val
			stack.PushBack(nodePathSumRight)
		}
	}
	return false
}
