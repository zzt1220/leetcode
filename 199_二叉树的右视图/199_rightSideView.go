package _99_二叉树的右视图

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
1、迭代，二叉树层序遍历
*/
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		var cur *TreeNode
		for i := 0; i < size; i++ {
			e := queue.Front()
			cur = e.Value.(*TreeNode)
			queue.Remove(e)
			if cur.Left != nil {
				queue.PushBack(cur.Left)
			}
			if cur.Right != nil {
				queue.PushBack(cur.Right)
			}
		}
		res = append(res, cur.Val)
	}
	return res
}
