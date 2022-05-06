package _05_从前序与中序遍历序列构造二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
1、递归（优化）
2、与第106道题类似
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	return traversal(preorder, 0, len(preorder), inorder, 0, len(inorder))
}

func traversal(preorder []int, preorderBegin int, preorderEnd int, inorder []int, inorderBegin int, inorderEnd int) *TreeNode {
	if preorderBegin == preorderEnd {
		return nil
	}
	rootValue := preorder[preorderBegin]
	root := new(TreeNode)
	root.Val = rootValue
	if preorderEnd == preorderBegin+1 {
		return root
	}
	var delimeterIndex int
	for delimeterIndex = inorderBegin; delimeterIndex < inorderEnd; delimeterIndex++ {
		if inorder[delimeterIndex] == rootValue {
			break
		}
	}
	leftInorderBegin := inorderBegin
	leftInorderEnd := delimeterIndex
	rightInorderBegin := delimeterIndex + 1
	rightInorderEnd := inorderEnd
	leftPerorderBegin := preorderBegin + 1
	preDelimeterIndex := preorderBegin + 1 + delimeterIndex - inorderBegin
	leftPreorderEnd := preDelimeterIndex
	rightPreorderBegin := preDelimeterIndex
	rightPreorderEnd := preorderEnd
	root.Left = traversal(preorder, leftPerorderBegin, leftPreorderEnd, inorder, leftInorderBegin, leftInorderEnd)
	root.Right = traversal(preorder, rightPreorderBegin, rightPreorderEnd, inorder, rightInorderBegin, rightInorderEnd)
	return root
}
