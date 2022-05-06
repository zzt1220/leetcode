package _00_二叉搜索树中的搜索

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

///*
//1、方法一：递归
//*/
//func searchBST(root *TreeNode, val int) *TreeNode {
//	if root == nil {
//		return nil
//	}
//	if val == root.Val {
//		return root
//	} else if val < root.Val {
//		return searchBST(root.Left, val)
//	} else {
//		return searchBST(root.Right, val)
//	}
//}

/*
1、方法二：迭代
*/
func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if val == root.Val {
			return root
		} else if val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}
