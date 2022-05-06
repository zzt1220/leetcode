package _06_从中序与后序遍历序列构造二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

///*
//1、递归(没优化)
//2、按照后续遍历的最后一个元素切割中序数组，再按照中序数组的切割下标切割后续数组
//*/
//func buildTree(inorder []int, postorder []int) *TreeNode {
//	if len(postorder) == 0 {
//		return nil
//	}
//	rootValue := postorder[len(postorder)-1]
//	root := new(TreeNode)
//	root.Val = rootValue
//	if len(postorder) == 1 {
//		return root
//	}
//	var delimeterIndex int
//	for delimeterIndex = 0; delimeterIndex < len(inorder); delimeterIndex++ {
//		if inorder[delimeterIndex] == rootValue {
//			break
//		}
//	}
//	leftInorder := inorder[:delimeterIndex]
//	rightInorder := inorder[delimeterIndex+1:]
//	leftPostorder := postorder[:delimeterIndex]
//	rightPostorder := postorder[delimeterIndex : len(postorder)-1]
//	root.Left = buildTree(leftInorder, leftPostorder)
//	root.Right = buildTree(rightInorder, rightPostorder)
//	return root
//}

/*
1、递归(优化)，没优化版本每次都需要拷贝切片，优化版本只需要指定区间，这里使用左闭右开方式
2、按照后续遍历的最后一个元素切割中序数组，再按照中序数组的切割下标切割后续数组
3、注意后续遍历数组的切割点和中序遍历切割点位置不同，后续遍历数组切割点是postDelimeterIndex:=postorderBegin+delimeterIndex-inorderBegin，这里很容易直接使用中序遍历的切割点
*/
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	return traversal(inorder, 0, len(inorder), postorder, 0, len(postorder))
}

func traversal(inorder []int, inorderBegin int, inorderEnd int, postorder []int, postorderBegin int, postorderEnd int) *TreeNode {
	if postorderEnd == postorderBegin {
		return nil
	}
	rootValue := postorder[postorderEnd-1]
	root := new(TreeNode)
	root.Val = rootValue
	if postorderEnd == postorderBegin+1 {
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
	leftPostorderBegin := postorderBegin
	postDelimeterIndex := postorderBegin + delimeterIndex - inorderBegin
	leftPostorderEnd := postDelimeterIndex
	rightPostorderBegin := postDelimeterIndex
	rightPostorderEnd := postorderEnd - 1
	root.Left = traversal(inorder, leftInorderBegin, leftInorderEnd, postorder, leftPostorderBegin, leftPostorderEnd)
	root.Right = traversal(inorder, rightInorderBegin, rightInorderEnd, postorder, rightPostorderBegin, rightPostorderEnd)
	return root
}
