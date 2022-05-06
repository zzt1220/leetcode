package _11_二叉树的最小深度

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

///*
//1、方法一：递归
//*/
//
//func minDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	if root.Left == nil && root.Right == nil {
//		return 1
//	}
//	depth := math.MaxInt32
//	if root.Left != nil {
//		depth = min(depth, minDepth(root.Left))
//	}
//	if root.Right != nil {
//		depth = min(depth, minDepth(root.Right))
//	}
//	return depth + 1
//}
//
//func min(x, y int) int {
//	if x > y {
//		return y
//	}
//	return x
//}

/*
1、方法二：迭代，层序遍历
*/

func minDepth(root *TreeNode) int {
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
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
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
