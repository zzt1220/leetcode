package _37_二叉树的层平均值

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
1、迭代，二叉树层序遍历
*/
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	res := make([]float64, 0)
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		total := 0
		for i := 0; i < size; i++ {
			e := queue.Front()
			cur := e.Value.(*TreeNode)
			queue.Remove(e)
			total += cur.Val
			if cur.Left != nil {
				queue.PushBack(cur.Left)
			}
			if cur.Right != nil {
				queue.PushBack(cur.Right)
			}
		}
		res = append(res, float64(total)/float64(size))
	}
	return res
}
