package striver_dsa_sheet

// Pblm: https://leetcode.com/problems/middle-of-the-linked-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	isEven := true
	for true {
		if fast.Next == nil {
			isEven = false
			break
		} else if fast.Next.Next == nil {

			break
		} else {
			fast = fast.Next.Next
			slow = slow.Next
		}

	}
	if isEven {
		return slow.Next
	}
	return slow

}
