package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

var TESTTIME map[int] string

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

func heapSortLinkList(head *ListNode) *ListNode  {
	heapData := createHeapList(head)
	dataLen := len(heapData)
	if dataLen == 0 {
		return nil
	} else if dataLen == 1 || dataLen == 2 {
		return head
	}
	newHead, heap := popNode(heapData)
	now := newHead
	for len(heap) > 0 {
		top, newHeap := popNode(heap)
		heap = newHeap
		now.Next = top
		now = now.Next
	}
	return newHead
}

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
		if data[0].Val < data[1].Val {
			return []*ListNode{data[0], data[1]}
		}
		return data
	} else if dataLen == 2 {
		if data[0].Val < node.Val {
			return []*ListNode{data[0], data[1], node}
		} else {
			return []*ListNode{node, data[0], data[1]}
		}
	}

	index := dataLen // (dataLen - 1) + 1, +1是因为append了一个
	data = append(data, node)
	data = adjustNodeBottom2Up(data, index)
	return data
}

func popNode(data []*ListNode) (*ListNode, []*ListNode) {
	dataLen := len(data)
	if len(data) == 0 {
		return nil, nil
	} else if dataLen == 1 {
		return data[0], nil
	} else if dataLen == 2 {
		return data[0], []*ListNode{data[1]}
	}

	top := data[0]
	data[0] = data[dataLen-1]
	data = data[:dataLen-1]
	data = adjustNodeTop2down(data, 0)
	return top, data
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

func adjustNodeTop2down(data []*ListNode, index int) []*ListNode {
	dataLen := len(data)
	node := data[index]
	leftChild := index * 2 + 1
	rightChild := index * 2 + 2
	for leftChild < dataLen {
		min := leftChild
		if rightChild < dataLen && data[leftChild].Val > data[rightChild].Val {
			min = rightChild
		}

		if data[index].Val < data[min].Val {
			break
		}
		data[index], data[min] = data[min], data[index]

		index = min
		leftChild = index * 2 + 1
		rightChild = index * 2 + 2
	}
	data[index] = node
	return data
}


/*
归并排序
*/
func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

func mergeSort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return merge(mergeSort(head, mid), mergeSort(mid, tail))
}

func mergeSortList(head *ListNode) *ListNode {
	return mergeSort(head, nil)
}


/*
* 共用方法
*/
func createListData(number int) []int {
	var data []int
	for i := 0; i < number; i++ {
		data = append(data, rand.Intn(number))
	}
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
		fmt.Print(" ", node.Val)
		if node.Next == nil {
			break
		}
		node = node.Next
	}
}

func checkResult(head, right *ListNode) string {
	now := head
	rightNow := right
	for now != nil {
		if now.Val != rightNow.Val {
			return "false"
		}
		now = now.Next
		rightNow = rightNow.Next
	}
	return "true"
}

func sortTestFunc(sortFunc func(head *ListNode) *ListNode, head, right *ListNode, name string) *ListNode {
	usHead := head
	fmt.Println("# "+ name +" start >>>>>>>>>>>>>>>>>>")
	start := time.Now().UnixNano()
	resData := sortFunc(usHead)
	total := time.Now().UnixNano() - start
	TESTTIME[int(total)] = name
	fmt.Println("# "+ name +" end, total time: ", total)
	fmt.Println("# "+ name +" sort result is ", checkResult(usHead, right))
	fmt.Println("# "+ name + " result: ")
	printListLink(resData)
	fmt.Println("\n")
	return resData
}

func main() {
	data := createListData(100)
	linkList := createListLink(data)
	fmt.Println("data >>>>>>>>>>>>>>>")
	printListLink(linkList)
	TESTTIME = make(map[int] string)
	retData := sortTestFunc(insertionSortList, linkList, &ListNode{}, "insertionSortList")

	linkList = createListLink(data)
	sortTestFunc(heapSortLinkList, linkList, retData, "heapSortLinkList")

	linkList = createListLink(data)
	sortTestFunc(mergeSortList, linkList, retData, "mergeSortList")

	var keyList []int
	for key := range TESTTIME {
		keyList = append(keyList, key)
	}
	sort.Ints(keyList)


	for i:=0; i<len(keyList); i++ {
		key := keyList[i]
		fmt.Println("time: ", key, "func name: ", TESTTIME[key])
	}
}