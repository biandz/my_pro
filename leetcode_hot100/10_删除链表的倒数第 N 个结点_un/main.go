package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := getNodeLength(head)
	v := l - n
	tl := 0
	for head != nil {
		if tl == v {
			tl++
			head = head.Next
			continue
		}
		tl++
		head = head.Next
	}

	return head

}

func getNodeLength(head *ListNode) int {
	l := 0
	for head != nil {
		l++
		head = head.Next
	}
	return l
}
