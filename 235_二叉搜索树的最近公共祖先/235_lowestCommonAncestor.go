package _35_二叉搜索树的最近公共祖先

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

///*
//1、方法一：递归
//*/
//func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//	if root.Val > p.Val && root.Val > q.Val {
//		return lowestCommonAncestor(root.Left, p, q)
//	} else if root.Val < p.Val && root.Val < q.Val {
//		return lowestCommonAncestor(root.Right, p, q)
//	} else {
//		return root
//	}
//}

/*
1、方法二：迭代
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}
