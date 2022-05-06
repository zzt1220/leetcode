package _45_二叉树的后序遍历

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

///*
//1、方法一：递归
//*/
//func postorderTraversal(root *TreeNode) []int {
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
//	traversal(cur.Right, res)
//	*res = append(*res, cur.Val)
//}

/*
1、方法二：迭代
2、与第144道题前序遍历类似，区别就是后续遍历是左子树先入栈，右子树后入栈，最后再将结果倒序输出
*/
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := list.New()
	if root == nil {
		return nil
	}
	stack.PushBack(root)
	for stack.Len() > 0 {
		e := stack.Back()
		node := e.Value.(*TreeNode)
		stack.Remove(e)

		res = append(res, node.Val)
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}
