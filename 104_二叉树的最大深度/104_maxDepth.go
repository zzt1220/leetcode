package _04_二叉树的最大深度

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

///*
//1、方法一：递归
//*/
//
//func maxDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
//}
//
//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}

/*
1、方法二：迭代，层序遍历
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		depth++
		size := queue.Len()
		for i := 0; i < size; i++ {
			e := queue.Front()
			cur := e.Value.(*TreeNode)
			queue.Remove(e)
			if cur.Left != nil {
				queue.PushBack(cur.Left)
			}
			if cur.Right != nil {
				queue.PushBack(cur.Right)
			}
		}
	}
	return depth
}
