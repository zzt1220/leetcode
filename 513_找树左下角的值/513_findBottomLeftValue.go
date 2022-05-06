package _13_找树左下角的值

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

///*
//1、方法一：递归，回溯
//*/
//var maxDepth int
//var maxLeftValue int
//
//func findBottomLeftValue(root *TreeNode) int {
//	maxDepth = -1
//	maxLeftValue = 0
//	if root == nil {
//		return 0
//	}
//	traversal(root, 0)
//	return maxLeftValue
//}
//
//func traversal(node *TreeNode, depth int) {
//	if node.Left == nil && node.Right == nil {
//		if depth > maxDepth {
//			maxDepth = depth
//			maxLeftValue = node.Val
//		}
//		return
//	}
//	if node.Left != nil {
//		traversal(node.Left, depth+1)
//	}
//	if node.Right != nil {
//		traversal(node.Right, depth+1)
//	}
//}

/*
1、方法二：层序遍历
*/
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxLeftValue := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			e := queue.Front()
			cur := e.Value.(*TreeNode)
			queue.Remove(e)
			if i == 0 {
				maxLeftValue = cur.Val
			}
			if cur.Left != nil {
				queue.PushBack(cur.Left)
			}
			if cur.Right != nil {
				queue.PushBack(cur.Right)
			}
		}
	}
	return maxLeftValue
}
