package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	current := l
	var res string
	for current != nil {
		res += strconv.Itoa(current.Val)
		if current.Next != nil {
			res += ","
		}
		current = current.Next
	}
	return res
}

func insertOrdered(list *ListNode, value int) {

	node := &ListNode{Val: value}
	if list.Next == nil {
		list.Next = node
		return
	}

	current := list.Next
	prev := list
	for current != nil {
		if value > current.Val && current.Next == nil {
			current.Next = node
			break
		}
		if value <= current.Val {
			prev.Next = node
			node.Next = current
			break
		}
		prev = current
		current = current.Next
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	res := &ListNode{}
	for _, list := range lists {
		listCurrent := list
		for listCurrent != nil {
			insertOrdered(res, listCurrent.Val)
			listCurrent = listCurrent.Next
		}
	}
	return res.Next
}

func main() {
	list1 := &ListNode{}
	list2 := &ListNode{}
	list3 := &ListNode{}

	initList := func(list *ListNode, values ...int) {
		current := list
		for i, value := range values {
			current.Val = value
			if i != len(values)-1 {
				current.Next = &ListNode{}
			}
			current = current.Next
		}
	}

	initList(list1, 1, 4, 5)
	initList(list2, 1, 3, 4)
	initList(list3, 2, 6)

	lists := []*ListNode{list1, list2, list3}
	fmt.Println(mergeKLists(lists))
}
