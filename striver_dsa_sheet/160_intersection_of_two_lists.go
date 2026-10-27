package striver_dsa_sheet

// https://leetcode.com/problems/intersection-of-two-linked-lists/description/

// For Intersection of Two Linked Lists, there are two common approaches:

// 1) Length Difference
// Find lengths of both lists.
// Move the longer list ahead by the length difference.
// Move both together until they meet.
// O(m+n) time, O(1) space
// 2) Two Pointers
// pA traverses A → B.
// pB traverses B → A.
// They automatically cover equal total distance and meet at the intersection.
// O(m+n) time, O(1) space
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA := headA
	pB := headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pB
}
