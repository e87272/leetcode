package main

import (
	"log"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	nowNode := removeNthFromEnd(makeTree([]int{1, 2}), 2)

	for {
		if nowNode == nil {
			break
		}
		log.Printf("removeNthFromEnd : %+v", nowNode.Val)
		if nowNode.Next != nil {
			nowNode = nowNode.Next
		} else {
			break
		}
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	_, countDown := fromEnd(head, n)
	if countDown == 0 {
		return head.Next
	}
	return head
}
func fromEnd(head *ListNode, n int) (*ListNode, int) {
	if head.Next == nil {
		return head, n - 1
	}
	nextNode, countDown := fromEnd(head.Next, n)
	if countDown == 0 {
		head.Next = nextNode.Next
	} else {
		head.Next = nextNode
	}
	return head, countDown - 1
}

func makeTree(ary []int) *ListNode {
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
