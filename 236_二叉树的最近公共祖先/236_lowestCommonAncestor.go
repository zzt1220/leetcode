package _36_二叉树的最近公共祖先

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
1、递归，后序遍历
2、本题关键在终止条件，一般后序遍历的终止条件都是判断节点指针是否为nil，本题需要新增判断节点值等于p或q时返回该节点指针
3、基于上述终止条件，当返回非nil节点时，说明该左或右子树中包含3种情况：
（1）值为p的节点
（2）值为q的节点
（3）值为p的节点和值为q的节点
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil && right == nil {
		return left
	}
	if left == nil && right != nil {
		return right
	}
	return nil
}
