package t2

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test_01(t *testing.T) {
	sli := []int{1, 2, 3, 4, 5}
	num := 2
	list := createNodeBySlice(sli)
	end := removeNthFromEnd(list, num)
	PrintNode(end)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sli := []int{}
	s := []int{}
	for head != nil {
		sli = append(sli, head.Val)
		head = head.Next
	}
	k := len(sli) - n
	s = append(s, sli[:k]...)
	s = append(s, sli[k+1:]...)
	return createNodeBySlice(s)
}

func createNodeBySlice(sli []int) *ListNode {
	var head, tail *ListNode
	for _, i := range sli {
		if head == nil {
			head = &ListNode{i, nil}
			tail = head
		} else {
			tail.Next = &ListNode{i, nil}
			tail = tail.Next
		}
	}
	return head
}

func PrintNode(list *ListNode) {
	//打印列表
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}

func Test_02(t *testing.T) {

}
