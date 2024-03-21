package t1

import (
	"container/list"
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test_01(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	rl := addTwoNumbers(l1, l2)
	rst := []int{}
	for {
		rst = append(rst, rl.Val)
		if rl.Next == nil {
			break
		}
		rl = rl.Next
	}
	fmt.Println(rst)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	var carry = 0

	for l1 != nil || l2 != nil {
		var n1, n2 = 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10

		if head == nil {
			head = &ListNode{sum, nil}
			tail = head
		} else {
			tail.Next = &ListNode{sum, nil}
			tail = tail.Next
		}
	}

	if carry > 0 {
		tail = &ListNode{carry, nil}
	}
	return head
}

func Test_02(t *testing.T) {
	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)

	for ele := myList.Front(); ele != nil; ele = ele.Next() {
		fmt.Println(ele.Value)
	}
}
