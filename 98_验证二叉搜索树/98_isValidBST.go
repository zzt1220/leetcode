package _8_验证二叉搜索树

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
//1、方法一：递归
//2、本质上就是中序遍历，中间节点要大于左边节点的最大值
//*/
//func isValidBST(root *TreeNode) bool {
//	maxValue := math.MinInt64
//	return traversal(root, &maxValue)
//}
//
//func traversal(root *TreeNode, maxValue *int) bool {
//	if root == nil {
//		return true
//	}
//	left := traversal(root.Left, maxValue)
//	if *maxValue < root.Val {
//		*maxValue = root.Val
//	} else {
//		return false
//	}
//	right := traversal(root.Right, maxValue)
//	return left && right
//}

/*
1、方法二：迭代
2、中序遍历的迭代法，稍微修改一下就行
*/
func isValidBST(root *TreeNode) bool {
	stack := list.New()
	if root == nil {
		return true
	}
	preValue := math.MinInt64
	cur := root
	for cur != nil || stack.Len() > 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			e := stack.Back()
			cur = e.Value.(*TreeNode)
			stack.Remove(e)
			if cur.Val > preValue {
				preValue = cur.Val
			} else {
				return false
			}
			cur = cur.Right
		}
	}
	return true
}
