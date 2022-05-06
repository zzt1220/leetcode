package _15_在每个树行中找最大值

/*
1、迭代，层序遍历
*/

import (
	"container/list"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		max := math.MinInt64
		for i := 0; i < size; i++ {
			e := queue.Front()
			cur := e.Value.(*TreeNode)
			queue.Remove(e)
			if cur.Val > max {
				max = cur.Val
			}
			if cur.Left != nil {
				queue.PushBack(cur.Left)
			}
			if cur.Right != nil {
				queue.PushBack(cur.Right)
			}
		}
		res = append(res, max)
	}
	return res
}
