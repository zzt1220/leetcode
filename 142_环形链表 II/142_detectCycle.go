package _42_环形链表_II

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
1、双指针
2、fast指针每次走两步，slow指针每次走一步，如果相遇就说明有环
3、从头节点和相遇节点开始走，再次相遇的节点就是环的入口节点
4、可以画图通过数学推导
*/
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}
