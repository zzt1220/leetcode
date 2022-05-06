package _04_左叶子之和

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

///*
//1、方法一：递归
//2、注意左叶子节点不是左侧节点，判断条件：root.Left != nil && root.Left.Left == nil && root.Left.Right == nil
//*/
//func sumOfLeftLeaves(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	mid := 0
//	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
//		mid = root.Left.Val
//	}
//	return mid + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
//}

/*
1、方法二：迭代
2、本质上就是使用前中后序遍历或层序遍历对每个节点进行判断
*/
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		e := stack.Back()
		node := e.Value.(*TreeNode)
		stack.Remove(e)
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			sum += node.Left.Val
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
	}
	return sum
}
