package _30_二叉搜索树的最小绝对差

import (
	"container/list"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

///*
//1、递归，中序遍历思想
//2、本题也可以直接使用中序遍历得到中序遍历结果，然后在对结果计算得到最小绝对差
//*/
//var minDiff int
//var preNode *TreeNode
//
//func getMinimumDifference(root *TreeNode) int {
//	minDiff = math.MaxInt64
//	preNode = nil
//	traversal(root)
//	return minDiff
//}
//
//func traversal(root *TreeNode) {
//	if root == nil {
//		return
//	}
//	traversal(root.Left)
//	if preNode != nil {
//		minDiff = min(minDiff, root.Val-preNode.Val)
//	}
//	preNode = root
//	traversal(root.Right)
//}
//
//func min(x, y int) int {
//	if x > y {
//		return y
//	}
//	return x
//}

/*
1、方法二：迭代
*/

func getMinimumDifference(root *TreeNode) int {
	minDiff := math.MaxInt64
	if root == nil {
		return minDiff
	}
	var preNode *TreeNode
	cur := root
	stack := list.New()
	for cur != nil || stack.Len() > 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			e := stack.Back()
			cur = e.Value.(*TreeNode)
			stack.Remove(e)
			if preNode != nil {
				minDiff = min(minDiff, cur.Val-preNode.Val)
			}
			preNode = cur
			cur = cur.Right
		}
	}
	return minDiff
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
