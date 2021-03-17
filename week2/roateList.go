package week2

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	before := new(ListNode)
	before = nil
	curr := head
	next := head.Next
	for next != nil {
		if next.Next == nil {
			curr.Next = before
			next.Next = curr
			return next
		}
		curr.Next = before
		before, curr, next = curr, next, next.Next
	}
	return nil
}
