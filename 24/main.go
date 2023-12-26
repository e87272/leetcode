package main

import "log"

func main() {
	list := swapPairs(makeLinkList([]int{1, 2, 3, 4}))
	for {
		if list == nil {
			break
		}
		log.Printf("mergeKLists : %+v", list.Val)
		if list.Next != nil {
			list = list.Next
		} else {
			break
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head != nil && head.Next != nil {
		temp := head.Next
		head.Next = temp.Next
		temp.Next = head
		head = temp
		if head.Next.Next != nil {
			subList := swapPairs(head.Next.Next)
			head.Next.Next = subList
		}
	}
	return head
}

func makeLinkList(ary []int) *ListNode {
	var nowNode *ListNode
	var firstNode *ListNode
	for k, v := range ary {
		newNode := &ListNode{
			Val:  v,
			Next: nil,
		}
		if k == 0 {
			firstNode = newNode
		} else {
			nowNode.Next = newNode
		}
		nowNode = newNode
	}
	return firstNode
}
