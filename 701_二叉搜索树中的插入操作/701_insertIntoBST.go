package _01_二叉搜索树中的插入操作

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

///*
//1、方法一：递归
//*/
//func insertIntoBST(root *TreeNode, val int) *TreeNode {
//	if root == nil {
//		node := new(TreeNode)
//		node.Val = val
//		return node
//	}
//	if root.Val > val {
//		root.Left = insertIntoBST(root.Left, val)
//	}
//	if root.Val < val {
//		root.Right = insertIntoBST(root.Right, val)
//	}
//	return root
//}

/*
1、方法二：迭代
2、迭代法需要一个parent指针
*/
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		node := new(TreeNode)
		node.Val = val
		return node
	}
	var parent *TreeNode
	cur := root
	for cur != nil {
		parent = cur
		if cur.Val > val {
			cur = cur.Left
		} else if cur.Val < val {
			cur = cur.Right
		}
	}
	node := new(TreeNode)
	node.Val = val
	if parent.Val > val {
		parent.Left = node
	}
	if parent.Val < val {
		parent.Right = node
	}
	return root
}
