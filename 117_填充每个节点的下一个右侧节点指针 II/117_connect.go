package _17_填充每个节点的下一个右侧节点指针_II

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/*
1、迭代，层序遍历
*/
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		tmp := queue
		queue = nil
		for i := 0; i < len(tmp); i++ {
			if i < len(tmp)-1 {
				tmp[i].Next = tmp[i+1]
			}
			if tmp[i].Left != nil {
				queue = append(queue, tmp[i].Left)
			}
			if tmp[i].Right != nil {
				queue = append(queue, tmp[i].Right)
			}
		}
	}
	return root
}
