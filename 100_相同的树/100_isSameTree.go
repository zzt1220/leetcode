package _00_相同的树

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

///*
//1、方法一：递归
//2、与第101道题思路一样
//*/
//func isSameTree(p *TreeNode, q *TreeNode) bool {
//	return compare(p, q)
//}
//
//func compare(pNode, qNode *TreeNode) bool {
//	if pNode == nil && qNode == nil {
//		return true
//	}
//	if pNode != nil && qNode == nil {
//		return false
//	}
//	if pNode == nil && qNode != nil {
//		return false
//	}
//	if pNode.Val != qNode.Val {
//		return false
//	}
//	return compare(pNode.Left, qNode.Left) && compare(pNode.Right, qNode.Right)
//}

/*
1、方法二：迭代
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	stack := list.New()
	stack.PushBack(p)
	stack.PushBack(q)
	for stack.Len() > 0 {
		e := stack.Front()
		pNode := e.Value.(*TreeNode)
		stack.Remove(e)
		e = stack.Front()
		qNode := e.Value.(*TreeNode)
		stack.Remove(e)
		if pNode == nil && qNode == nil {
			continue
		}
		if pNode != nil && qNode == nil {
			return false
		}
		if pNode == nil && qNode != nil {
			return false
		}
		if pNode.Val != qNode.Val {
			return false
		}
		stack.PushBack(pNode.Left)
		stack.PushBack(qNode.Left)
		stack.PushBack(pNode.Right)
		stack.PushBack(qNode.Right)
	}
	return true
}
