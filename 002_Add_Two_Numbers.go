package go_leetcode
type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := l1.Val + l2.Val
	carry := sum / 10
	ret := &ListNode{sum % 10, nil}
	p1 := l1
	p2 := l2
	p3 := ret
	for p1.Next!=nil && p2.Next!=nil{
		sum = carry + p1.Next.Val + p2.Next.Val
		node := &ListNode{sum % 10, nil}
		carry = sum / 10
		p3.Next = node
		p1 = p1.Next
		p2 = p2.Next
		p3 = p3.Next
	}
	for p1.Next!=nil{
		sum = carry + p1.Next.Val
		node := &ListNode{sum % 10, nil}
		carry = sum / 10
		p3.Next = node
		p1 = p1.Next
		p3 = p3.Next
	}
	for p2.Next!=nil{
		sum = carry + p2.Next.Val
		node := &ListNode{sum % 10, nil}
		carry = sum / 10
		p3.Next = node
		p2 = p2.Next
		p3 = p3.Next
	}
	if carry != 0{
		node := &ListNode{carry, nil}
		p3.Next = node
	}
	return ret
}