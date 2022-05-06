package _17_合并二叉树

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

///*
//1、方法一：递归
//*/
//func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
//	if root1 == nil && root2 == nil {
//		return nil
//	}
//	if root1 != nil && root2 == nil {
//		return root1
//	}
//	if root1 == nil && root2 != nil {
//		return root2
//	}
//	root1.Val = root1.Val + root2.Val
//	root1.Left = mergeTrees(root1.Left, root2.Left)
//	root1.Right = mergeTrees(root1.Right, root2.Right)
//	return root1
//}

/*
1、方法二：迭代
*/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	queue := list.New()
	queue.PushBack(root1)
	queue.PushBack(root2)
	for queue.Len() > 0 {
		e1 := queue.Front()
		node1 := e1.Value.(*TreeNode)
		queue.Remove(e1)
		e2 := queue.Front()
		node2 := e2.Value.(*TreeNode)
		queue.Remove(e2)
		node1.Val = node1.Val + node2.Val
		if node1.Left != nil && node2.Left != nil {
			queue.PushBack(node1.Left)
			queue.PushBack(node2.Left)
		}
		if node1.Right != nil && node2.Right != nil {
			queue.PushBack(node1.Right)
			queue.PushBack(node2.Right)
		}
		if node1.Left == nil && node2.Left != nil {
			node1.Left = node2.Left
		}
		if node1.Right == nil && node2.Right != nil {
			node1.Right = node2.Right
		}
	}
	return root1
}
