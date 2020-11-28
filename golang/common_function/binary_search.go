package main

import (
	"fmt"
	"sort"
)

func binarySearch(data []int, index int) int {
	dataLen := len(data)
	if dataLen == 0 {
		return -1
	}
	// 如果数据量比较小，那么不需要用二分查找也可
	if dataLen < 10 {
		for i:=0; i<dataLen; i++ {
			if data[i] == index {
				return i
			}
		}
		return -1
	}

	left, right := 0, dataLen - 1
	for right > left {
		flag := int(left + (right - left) / 2)

		if data[flag] > index {
			right = flag
		} else if data[flag] < index {
			left = flag + 1
		} else {
			right--
		}
	}
	return left
}

func main() {
	data := []int{1, 5, 15, 74, 82, 16, 94, 13, 15, 12, 15, 22, 36, 54, 78, 66, 1, 16, 3, 1, 2, 0, 42}
	sort.Ints(data)
	fmt.Println(data)
	fmt.Println(binarySearch(data, 1))
}
