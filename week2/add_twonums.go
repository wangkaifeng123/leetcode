package week2

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	curr := dummy
	curr.Val = 0

	var temp int
	for l1 != nil || l2 != nil || temp != 0 {
		sum := 0
		curr.Next = new(ListNode)
		curr = curr.Next
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += temp
		curr.Val = sum % 10
		temp = sum / 10
	}
	return dummy.Next
}
