package _44_二叉树的前序遍历

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
//func preorderTraversal(root *TreeNode) []int {
//	res := make([]int, 0)
//	traversal(root, &res)
//	return res
//}
//
//func traversal(cur *TreeNode, res *[]int) {
//	if cur == nil {
//		return
//	}
//	*res = append(*res, cur.Val)
//	traversal(cur.Left, res)
//	traversal(cur.Right, res)
//}

/*
1、方法二：迭代
2、注意list的使用方法
3、题解中入栈的都是非nil节点，所以不需要使用node != nil这个判断条件
4、先右子树入栈，再左子树入栈
*/
func preorderTraversal(root *TreeNode) []int {
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
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return res
}
