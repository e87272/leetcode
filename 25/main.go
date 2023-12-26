package main

import "log"

func main() {
	// reverseKGroup(makeLinkList([]int{1, 2, 3, 4, 5}), 5)
	list := reverseKGroup(makeLinkList([]int{1, 2, 3, 4, 5}), 4)
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

func reverseKGroup(head *ListNode, k int) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	nowNode := head
	for i := 0; i < k-1; i++ {
		nowNode = nowNode.Next
		if nowNode == nil {
			return head
		}
	}
	var subList *ListNode = nil
	if nowNode.Next != nil {
		subList = reverseKGroup(nowNode.Next, k)
	}
	reverseLinkK(head, k)
	head.Next = subList
	return nowNode
}

func reverseLinkK(head *ListNode, k int) *ListNode {
	if k == 1 {
		// log.Printf("reverseLinkK head : %+v", head)
		return head
	}
	nowNode := reverseLinkK(head.Next, k-1)
	nowNode.Next = head
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
