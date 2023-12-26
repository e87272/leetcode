package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := mergeTwoLists(makeLinkList([]int{1, 2, 4}), makeLinkList([]int{1, 3, 4}))
	for {
		if list == nil {
			break
		}
		log.Printf("mergeTwoLists : %+v", list.Val)
		if list.Next != nil {
			list = list.Next
		} else {
			break
		}
	}
}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var firstNode *ListNode
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		firstNode = list1
		list1 = list1.Next
	} else {
		firstNode = list2
		list2 = list2.Next
	}
	nowNode := firstNode
	for {
		if list1 == nil && list2 == nil {
			break
		} else if list1 == nil {
			nowNode.Next = list2
			break
		} else if list2 == nil {
			nowNode.Next = list1
			break
		} else {
			if list1.Val < list2.Val {
				nowNode.Next = list1
				list1 = list1.Next
			} else {
				nowNode.Next = list2
				list2 = list2.Next
			}
			nowNode = nowNode.Next
		}
	}
	return firstNode
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else {
		if list1.Val < list2.Val {
			list1.Next = mergeTwoLists(list1.Next, list2)
			return list1
		} else {
			list2.Next = mergeTwoLists(list1, list2.Next)
			return list2
		}
	}
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
