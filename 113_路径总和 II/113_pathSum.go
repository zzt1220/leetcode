package _13_路径总和_II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
1、方法一：递归
*/
var res [][]int

func pathSum(root *TreeNode, targetSum int) [][]int {
	res = [][]int{}
	var path []int
	if root == nil {
		return res
	}
	traversal(root, targetSum, 0, path)
	return res
}

func traversal(root *TreeNode, targetSum int, curSum int, path []int) {
	curSum += root.Val
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		if curSum == targetSum {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
	}
	if root.Left != nil {
		traversal(root.Left, targetSum, curSum, path)
	}
	if root.Right != nil {
		traversal(root.Right, targetSum, curSum, path)
	}
}

///*
//1、方法二：迭代
//2、以下迭代法耗时间和空间，本题建议使用递归解法
//3、这里迭代法本质上市先序遍历，也可以中序，后续，层序遍历
//*/
//type NodePathSum struct {
//	Node *TreeNode
//	Sum  int
//	Path []int
//}
//
//var res [][]int
//
//func pathSum(root *TreeNode, targetSum int) [][]int {
//	res = [][]int{}
//	if root == nil {
//		return res
//	}
//	stack := list.New()
//	rootSum := new(NodePathSum)
//	rootSum.Node = root
//	rootSum.Sum = root.Val
//	rootSum.Path = append(rootSum.Path, root.Val)
//	stack.PushBack(rootSum)
//	for stack.Len() > 0 {
//		e := stack.Back()
//		nodePathSum := e.Value.(*NodePathSum)
//		stack.Remove(e)
//		if nodePathSum.Node.Left == nil && nodePathSum.Node.Right == nil && nodePathSum.Sum == targetSum {
//			tmp := make([]int, len(nodePathSum.Path))
//			copy(tmp, nodePathSum.Path)
//			res = append(res, tmp)
//			continue
//		}
//		if nodePathSum.Node.Left != nil {
//			nodePathSumLeft := new(NodePathSum)
//			nodePathSumLeft.Node = nodePathSum.Node.Left
//			nodePathSumLeft.Sum = nodePathSum.Sum + nodePathSum.Node.Left.Val
//			nodePathSumLeft.Path = make([]int, len(nodePathSum.Path))
//			copy(nodePathSumLeft.Path, nodePathSum.Path)
//			nodePathSumLeft.Path = append(nodePathSumLeft.Path, nodePathSum.Node.Left.Val)
//			stack.PushBack(nodePathSumLeft)
//		}
//		if nodePathSum.Node.Right != nil {
//			nodePathSumRight := new(NodePathSum)
//			nodePathSumRight.Node = nodePathSum.Node.Right
//			nodePathSumRight.Sum = nodePathSum.Sum + nodePathSum.Node.Right.Val
//			nodePathSumRight.Path = make([]int, len(nodePathSum.Path))
//			copy(nodePathSumRight.Path, nodePathSum.Path)
//			nodePathSumRight.Path = append(nodePathSumRight.Path, nodePathSum.Node.Right.Val)
//			stack.PushBack(nodePathSumRight)
//		}
//	}
//	return res
//}
