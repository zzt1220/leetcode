package _4_两两交换链表中的节点

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
1、pre指向上两个节点的最后一个节点，head指向下两个节点的第一个节点
2、使用虚拟头节点
*/
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	pre := dummy
	for head != nil && head.Next != nil {
		pre.Next = head.Next
		next := head.Next.Next
		head.Next.Next = head
		head.Next = next

		pre = head
		head = next
	}
	return dummy.Next
}
