package _38_把二叉搜索树转换为累加树

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

///*
//1、方法一：递归
//2、使用反中序遍历，使用一个变量记录前一个累加值
//*/
//func convertBST(root *TreeNode) *TreeNode {
//	pre := 0
//	traversal(root, &pre)
//	return root
//}
//
//func traversal(cur *TreeNode, pre *int) {
//	if cur == nil {
//		return
//	}
//	traversal(cur.Right, pre)
//	cur.Val = cur.Val + *pre
//	*pre = cur.Val
//	traversal(cur.Left, pre)
//}

/*
1、方法二：迭代
2、反中序遍历思想
*/
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	pre := 0
	stack := list.New()
	cur := root
	for cur != nil || stack.Len() > 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Right
		} else {
			e := stack.Back()
			cur = e.Value.(*TreeNode)
			stack.Remove(e)
			cur.Val = cur.Val + pre
			pre = cur.Val
			cur = cur.Left
		}
	}
	return root
}
