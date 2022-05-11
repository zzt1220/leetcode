package _06_反转链表

type ListNode struct {
	Val  int
	Next *ListNode
}

///*
//1、方法一：递归
//2、从后往前翻转
//*/
//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	r := reverseList(head.Next)
//	head.Next.Next = head
//	head.Next = nil
//	return r
//}

/*
1、方法二：双指针
2、初始化pre为nil指针
3、从前往后翻转
*/
func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode
	pre = nil
	cur = head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
