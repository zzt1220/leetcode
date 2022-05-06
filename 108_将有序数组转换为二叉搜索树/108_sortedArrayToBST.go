package _08_将有序数组转换为二叉搜索树

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

///*
//1、方法一：递归
//*/
//func sortedArrayToBST(nums []int) *TreeNode {
//	return traversal(nums, 0, len(nums)-1)
//}
//func traversal(nums []int, left int, right int) *TreeNode {
//	if left > right {
//		return nil
//	}
//	mid := (left + right) / 2
//	node := new(TreeNode)
//	node.Val = nums[mid]
//	node.Left = traversal(nums, left, mid-1)
//	node.Right = traversal(nums, mid+1, right)
//	return node
//}

/*
1、方法二：迭代
2、迭代需要把节点、左区间、右区间一起入栈
*/

type NodeRange struct {
	Node  *TreeNode
	Left  int
	Right int
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	nodeRange := new(NodeRange)
	root := new(TreeNode)
	nodeRange.Node = root
	nodeRange.Left = 0
	nodeRange.Right = len(nums) - 1
	nodeRangeStack := list.New()
	nodeRangeStack.PushBack(nodeRange)
	for nodeRangeStack.Len() > 0 {
		e := nodeRangeStack.Back()
		nodeRange := e.Value.(*NodeRange)
		nodeRangeStack.Remove(e)
		mid := (nodeRange.Left + nodeRange.Right) / 2
		nodeRange.Node.Val = nums[mid]
		if nodeRange.Left <= mid-1 {
			nodeRangeL := new(NodeRange)
			node := new(TreeNode)
			nodeRangeL.Node = node
			nodeRangeL.Left = nodeRange.Left
			nodeRangeL.Right = mid - 1
			nodeRangeStack.PushBack(nodeRangeL)
			nodeRange.Node.Left = nodeRangeL.Node
		}
		if nodeRange.Right >= mid+1 {
			nodeRangeR := new(NodeRange)
			node := new(TreeNode)
			nodeRangeR.Node = node
			nodeRangeR.Left = mid + 1
			nodeRangeR.Right = nodeRange.Right
			nodeRangeStack.PushBack(nodeRangeR)
			nodeRange.Node.Right = nodeRangeR.Node
		}
	}
	return nodeRange.Node
}
