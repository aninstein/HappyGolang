package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

/*
题目：链表插入排序
题目链接：https://leetcode-cn.com/problems/insertion-sort-list/

输入: 4->2->1->3
输出: 1->2->3->4

*/
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


/*
题目：对链表进行排序
题目链接：https://leetcode-cn.com/problems/sort-list/

输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]

*/
func createHeapList(head *ListNode) []*ListNode {
	if head == nil {
		return nil
	}
	var heapData []*ListNode
	now := head
	for now != nil {
		heapData = insertMinHeap(heapData, now)
		now = now.Next
	}
	return heapData
}

func insertMinHeap(data []*ListNode, node *ListNode) []*ListNode {
	if data == nil {
		return []*ListNode{node}
	}
	dataLen := len(data)
	if dataLen == 1 {
		data = append(data, node)
		if data[0].Val > data[1].Val {
			return []*ListNode{data[0], data[1]}
		}
		return data
	} else if dataLen == 2 {
		if data[0].Val > node.Val {
			return []*ListNode{data[0], data[1], node}
		} else {
			return []*ListNode{node, data[0], data[1]}
		}
	}

	index := dataLen // (dataLen - 1) + 1, +1是因为append了一个
	data = append(data, node)
	adjustNodeBottom2Up(data, index)
	return data
}

func adjustNodeBottom2Up(data []*ListNode, index int) []*ListNode {
	parent := int((index - 1) / 2)
	node := data[index]
	for parent >= 0 {
		if data[index].Val > data[parent].Val {
			break
		}

		data[parent], data[index] = data[index], data[parent]
		index = parent
		parent = int((index - 1) / 2)
	}
	data[index] = node
	return data
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