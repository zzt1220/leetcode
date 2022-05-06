package _01_对称二叉树

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

///*
//1、方法一：递归
//2、判断是否对称需要比较外侧节点与外侧节点比较，内侧节点与内侧节点比较
//*/
//func isSymmetric(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	return compare(root.Left, root.Right)
//}
//
//func compare(left, right *TreeNode) bool {
//	if left == nil && right == nil {
//		return true
//	}
//	if left != nil && right == nil {
//		return false
//	}
//	if left == nil && right != nil {
//		return false
//	}
//	if left.Val != right.Val {
//		return false
//	}
//	return compare(left.Left, right.Right) && compare(left.Right, right.Left)
//}

/*
1、方法二：迭代
2、判断是否对称需要比较外侧节点与外侧节点比较，内侧节点与内侧节点比较
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := list.New()
	stack.PushBack(root.Left)
	stack.PushBack(root.Right)
	for stack.Len() > 0 {
		e := stack.Front()
		left := e.Value.(*TreeNode)
		stack.Remove(e)
		e = stack.Front()
		right := e.Value.(*TreeNode)
		stack.Remove(e)
		if left == nil && right == nil {
			continue
		}
		if left != nil && right == nil {
			return false
		}
		if left == nil && right != nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		stack.PushBack(left.Left)
		stack.PushBack(right.Right)
		stack.PushBack(left.Right)
		stack.PushBack(right.Left)
	}
	return true
}
