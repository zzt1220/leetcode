package _50_删除二叉搜索树中的节点

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

///*
//1、方法一：递归（正常使用这个解法，方便理解）
//2、删除节点有5中情况：
//（1）找不到删除的节点
//（2）删除的节点左子树和右子树都是空
//（3）删除的节点左子树空，右子树非空
//（4）删除的节点左子树非空，右子树空
//（5）删除的节点左子树和右子树都非空
//*/
//func deleteNode(root *TreeNode, key int) *TreeNode {
//	if root == nil {
//		return nil
//	}
//	if root.Val == key {
//		if root.Left == nil && root.Right == nil {
//			return nil
//		}
//		if root.Left == nil {
//			return root.Right
//		}
//		if root.Right == nil {
//			return root.Left
//		}
//		cur := root.Right
//		for cur.Left != nil {
//			cur = cur.Left
//		}
//		cur.Left = root.Left
//		return root.Right
//	}
//	if root.Val > key {
//		root.Left = deleteNode(root.Left, key)
//	}
//	if root.Val < key {
//		root.Right = deleteNode(root.Right, key)
//	}
//	return root
//}

/*
1、方法二：递归
2、找到要删除的节点，然后跟该节点的右子树最左侧节点交换数值，递归多次交换后，会出现删除值的节点右子树为空的情况，会被删除掉
3、注意不能使用判断条件，因为root的值可能被交换过
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
*/
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Right == nil {
			return root.Left
		} else {
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			root.Val, cur.Val = cur.Val, root.Val
		}
	}
	root.Left = deleteNode(root.Left, key)
	root.Right = deleteNode(root.Right, key)
	return root
}
