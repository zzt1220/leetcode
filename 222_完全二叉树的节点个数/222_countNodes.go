package _22_完全二叉树的节点个数

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

///*
//1、方法一：递归
//*/
//func countNodes(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	return countNodes(root.Left) + countNodes(root.Right) + 1
//}
//
///*
//1、方法二：迭代，使用层序遍历方法计数，也可以使用前中后序遍历
//*/
//func countNodes(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	queue := list.New()
//	queue.PushBack(root)
//	count := 0
//	for queue.Len() > 0 {
//		size := queue.Len()
//		for i := 0; i < size; i++ {
//			count++
//			e := queue.Front()
//			cur := e.Value.(*TreeNode)
//			queue.Remove(e)
//			if cur.Left != nil {
//				queue.PushBack(cur.Left)
//			}
//			if cur.Right != nil {
//				queue.PushBack(cur.Right)
//			}
//		}
//	}
//	return count
//}

/*
1、方法三：使用完全二叉树的性质
*/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := root.Left
	right := root.Right
	leftH := 0
	rightH := 0
	for left != nil {
		left = left.Left
		leftH++
	}
	for right != nil {
		right = right.Right
		rightH++
	}
	if leftH == rightH {
		return 2<<leftH - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
