package _57_二叉树的所有路径

import (
	"strconv"
)

/*
1、递归回溯算法
2、注意go语言的切片本质上包含指针，长度，容量三个元素都是值拷贝
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []string

func binaryTreePaths(root *TreeNode) []string {
	res = []string{}
	var path []int
	if root == nil {
		return res
	}
	traversal(root, path)
	return res
}

func traversal(node *TreeNode, path []int) {
	path = append(path, node.Val)
	if node.Left == nil && node.Right == nil {
		var sPath string
		for i := 0; i < len(path)-1; i++ {
			sPath = sPath + strconv.Itoa(path[i]) + "->"
		}
		sPath = sPath + strconv.Itoa(path[len(path)-1])
		res = append(res, sPath)
		return
	}
	if node.Left != nil {
		traversal(node.Left, path)
	}
	if node.Right != nil {
		traversal(node.Right, path)
	}
}
