package _01_二叉搜索树中的众数

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

///*
//1、递归，使用中序遍历思想
//2、也可以直接使用中序遍历出结果，然后对结果做统计计算得到众数
//*/
//var count int
//var maxCount int
//var preNode *TreeNode
//var res []int
//
//func findMode(root *TreeNode) []int {
//	count = 0
//	maxCount = 0
//	preNode = nil
//	res = []int{}
//	traversal(root)
//	return res
//}
//
//func traversal(root *TreeNode) {
//	if root == nil {
//		return
//	}
//	traversal(root.Left)
//	if preNode == nil {
//		count = 1
//	} else if preNode.Val == root.Val {
//		count++
//	} else {
//		count = 1
//	}
//	if count == maxCount {
//		res = append(res, root.Val)
//	}
//	if count > maxCount {
//		maxCount = count
//		res = []int{}
//		res = append(res, root.Val)
//	}
//	preNode = root
//	traversal(root.Right)
//}

/*
1、迭代，与中序遍历代码类似
*/
func findMode(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	count := 0
	maxCount := 0
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
			if preNode == nil {
				count = 1
			} else if preNode.Val == cur.Val {
				count++
			} else {
				count = 1
			}
			if count == maxCount {
				res = append(res, cur.Val)
			}
			if count > maxCount {
				maxCount = count
				res = []int{}
				res = append(res, cur.Val)
			}
			preNode = cur
			cur = cur.Right
		}
	}
	return res
}
