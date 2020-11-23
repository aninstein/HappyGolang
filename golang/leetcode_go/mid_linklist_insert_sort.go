package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	now := head
	newHead := now
	indexNode := newHead  // 因为这个链表没有pre指针，所以每一次插入都是从头开始
	for now.Next != nil {
		if now.Val <= now.Next.Val {
			now = now.Next
			continue
		}
		insertNode := now.Next
		var preNode *ListNode
		for true {
			if indexNode.Val >= insertNode.Val {
				now.Next = insertNode.Next
				if preNode != nil {
					preNode.Next = insertNode
				} else {
					newHead = insertNode
				}
				insertNode.Next = indexNode
				indexNode = newHead  // 重置index
				break
			} else {
				preNode = indexNode
				indexNode = indexNode.Next
			}
		}
	}
	return newHead
}

func createListLink(data []int) *ListNode {
	dataLen := len(data)
	if dataLen == 0 {
		return nil
	}
	head := &ListNode{Val: data[0]}
	if dataLen == 1 {
		return head
	}

	now := head
	first := now
	for i:=1; i<dataLen; i++ {
		now.Next = &ListNode{Val: data[i]}
		now = now.Next
	}
	return first
}

func printListLink(node *ListNode)  {
	for true {
		fmt.Println(node.Val)
		if node.Next == nil {
			break
		}
		node = node.Next
	}
}

func main() {
	data := []int{1, 5, 8, 5, 4, 2, 3, 4, 7, 9}
	res := createListLink(data)
	fmt.Println("data >>>>>>>>>>>>>>>")
	printListLink(res)
	fmt.Println("sort data >>>>>>>>>>>>>")
	sortRes := insertionSortList(res)
	printListLink(sortRes)
}