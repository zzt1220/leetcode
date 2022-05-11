package _9_删除链表的倒数第_N_个结点

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
1、双指针
2、fast指针比slow指针快n+1步
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head
	fast := dummyHead
	slow := dummyHead
	for n >= 0 && fast != nil {
		fast = fast.Next
		n--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}
