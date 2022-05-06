package _29_N_叉树的层序遍历

/*
1、迭代，层序遍历
*/

import "container/list"

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
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
			cur := e.Value.(*Node)
			queue.Remove(e)
			levelRes[i] = cur.Val
			for j := 0; j < len(cur.Children); j++ {
				queue.PushBack(cur.Children[j])
			}
		}
		res = append(res, levelRes)
	}
	return res
}
