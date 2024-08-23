package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func (head *ListNode) Insert(val int, pos int) *ListNode {
	if pos == 0 {
		return &ListNode{Val: val, Next: head}
	}
	var dummy *ListNode = head

	for i := 0; i < pos-1 && dummy.Next != nil; i++ {
		dummy = dummy.Next
	}
	temp := ListNode{}
	temp.Val = val
	temp.Next = dummy.Next
	dummy.Next = &temp
	return head
}


func PrintList(head *ListNode) {
	dummy := head
	for dummy != nil {
		fmt.Println(dummy.Val)
		dummy = dummy.Next
	}
}

func main() {
	x := NListNode(1)
	x.Next = NListNode(2)
	x.Next.Next = NListNode(3)
	fmt.Println("FIRST LIST:")
	PrintList(x)
	head := x.Insert(34343, 1)
	fmt.Println("SECOND LIST:")
	PrintList(head)
}
