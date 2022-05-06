package _59_N_叉树的最大深度

import "container/list"

type Node struct {
	Val      int
	Children []*Node
}

///*
//1、方法一：递归
//2、递归方法中+1操作是关键
// */
//
//func maxDepth(root *Node) int {
//	if root == nil {
//		return 0
//	}
//	depth := 0
//	for i := 0; i < len(root.Children); i++ {
//		depthtmp := maxDepth(root.Children[i])
//		if depthtmp > depth {
//			depth = depthtmp
//		}
//	}
//	return depth + 1
//}

/*
1、方法一：迭代
*/
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		depth++
		size := queue.Len()
		for i := 0; i < size; i++ {
			e := queue.Front()
			cur := e.Value.(*Node)
			queue.Remove(e)
			for j := 0; j < len(cur.Children); j++ {
				if cur.Children[j] != nil {
					queue.PushBack(cur.Children[j])
				}
			}
		}
	}
	return depth
}
