package main

import (
	"log"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := mergeKLists([]*ListNode{makeLinkList([]int{1, 2, 4}), makeLinkList([]int{1, 3, 4})})
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

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var firstNode *ListNode
	var min int = math.MaxInt32
	var nodeIndex int
	for k, v := range lists {
		if v != nil && v.Val < min {
			min = v.Val
			firstNode = v
			nodeIndex = k
		}
	}
	if min == math.MaxInt32 {
		return nil
	}
	lists[nodeIndex] = lists[nodeIndex].Next
	nowNode := firstNode
	for {
		min = math.MaxInt32
		for k, v := range lists {
			if v != nil {
				if v.Val < min {
					min = v.Val
					nodeIndex = k
				}
			}
		}
		if min == math.MaxInt32 {
			break
		}
		nowNode.Next = lists[nodeIndex]
		nowNode = nowNode.Next
		lists[nodeIndex] = lists[nodeIndex].Next
	}
	return firstNode
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
