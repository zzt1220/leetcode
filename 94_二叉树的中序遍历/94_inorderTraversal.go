package _4_二叉树的中序遍历

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

///*
//1、方法一：递归
//2、注意切片
//*/
//func inorderTraversal(root *TreeNode) []int {
//	res := make([]int, 0)
//	traversal(root, &res)
//	return res
//}
//
//func traversal(cur *TreeNode, res *[]int) {
//	if cur == nil {
//		return
//	}
//	traversal(cur.Left, res)
//	*res = append(*res, cur.Val)
//	traversal(cur.Right, res)
//}

/*
1、方法二：迭代
2、与前序遍历和后续遍历不同的是需要一个cur表示当前节点，相比之下比前序遍历和中序遍历的迭代法难一些
*/
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := list.New()
	if root == nil {
		return nil
	}
	cur := root
	for cur != nil || stack.Len() > 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			e := stack.Back()
			cur = e.Value.(*TreeNode)
			stack.Remove(e)
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}
