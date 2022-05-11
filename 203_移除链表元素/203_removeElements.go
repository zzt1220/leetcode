package _03_移除链表元素

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
1、使用虚拟头节点指向链表头结点，方便程序处理
*/
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
