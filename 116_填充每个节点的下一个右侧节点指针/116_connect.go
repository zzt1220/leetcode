package _16_填充每个节点的下一个右侧节点指针

/*
1、迭代，层序遍历
*/

import "container/list"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		var pre *Node
		var cur *Node
		for i := 0; i < size; i++ {
			e := queue.Front()
			queue.Remove(e)
			if i == 0 {
				cur = e.Value.(*Node)
				pre = cur
			} else {
				cur = e.Value.(*Node)
				pre.Next = cur
				pre = cur
			}
			if cur.Left != nil {
				queue.PushBack(cur.Left)
			}
			if cur.Right != nil {
				queue.PushBack(cur.Right)
			}
		}
		cur.Next = nil
	}
	return root
}
