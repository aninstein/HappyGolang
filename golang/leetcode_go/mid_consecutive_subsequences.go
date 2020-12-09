package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
* 题目：分割数组为连续子序列
* 题目链接：https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences/
* 输入: [1,2,3,3,4,4,5,5]
输出: True
解释:
你可以分割出这样两个连续子序列 :
1, 2, 3, 4, 5
3, 4, 5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/


//func isPossible(nums []int) bool {
//	dataLen := len(nums)
//	if dataLen == 0 {
//		return true
//	} else if dataLen == 1 || dataLen == 2 {
//		return false
//	}
//	var exNums []int
//	nowNum := nums[0]
//	needNum := nowNum + 1
//	fmt.Print(" ", nowNum)
//	for i := 1; i < len(nums); i++ {
//		num := nums[i]
//		if num != needNum {
//			if num != nowNum {
//				exNums = append(exNums, nums[i:]...)
//				break
//			} else {
//				exNums = append(exNums, num)
//				continue
//			}
//		} else {
//			nowNum = num
//			needNum = nowNum + 1
//			fmt.Print(" ", nowNum)
//		}
//	}
//	fmt.Println("\n")
//	return isPossible(exNums)
//}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp) push(v int)         { heap.Push(h, v) }
func (h *hp) pop() int           { return heap.Pop(h).(int) }

func isPossible(nums []int) bool {
	lens := map[int]*hp{}
	for _, v := range nums {
		if lens[v] == nil {
			lens[v] = new(hp)
		}
		if h := lens[v-1]; h != nil {
			prevLen := h.pop()
			if h.Len() == 0 {
				delete(lens, v-1)
			}
			lens[v].push(prevLen + 1)
		} else {
			lens[v].push(1)
		}
	}

	for _, h := range lens {
		if h.IntSlice[0] < 3 {
			return false
		}
	}
	return true
}

func main() {
	data := []int{1,2,3,3,4,4,5,5}
	fmt.Println(isPossible(data))
}