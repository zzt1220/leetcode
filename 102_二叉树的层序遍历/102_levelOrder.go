package _02_二叉树的层序遍历

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
1、迭代
2、与前序遍历迭代法的区别是使用队列先进先出，先进左子树后进右子树
*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		levelRes := make([]int, size)
		for i := 0; i < size; i++ {
			e := queue.Front()
			cur := e.Value.(*TreeNode)
			queue.Remove(e)
			levelRes[i] = cur.Val
			if cur.Left != nil {
				queue.PushBack(cur.Left)
			}
			if cur.Right != nil {
				queue.PushBack(cur.Right)
			}
		}
		res = append(res, levelRes)
	}
	return res
}
