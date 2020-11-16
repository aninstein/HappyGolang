package main

import (
	"fmt"
	"sort"
)

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	index := []int{r0, c0}
	dataMap := make(map[int][][]int)
	for i:=0; i<R; i++ {
		for j:=0; i<C; j++ {
			var cell []int
			cell = append(cell, i)
			cell = append(cell, j)

			dis := manhatan(index, cell)
			_, ok := dataMap[dis]
			if ok {
				dataMap[dis] = append(dataMap[dis], cell)
			} else {
				dataMap[dis] = [][]int{cell}
			}
		}
	}

	var mapKeyList []int
	for k, _ := range dataMap {
		mapKeyList = append(mapKeyList,k)
	}
	sort.Ints(mapKeyList)

	var mt [][]int
	for i:=0; i<len(mapKeyList); i++ {
		mt = append(mt, dataMap[i]...)
	}
	return mt
}

func manhatan(a, b []int) int {
	return abs(a[0] - b[0]) + abs(a[1] - b[1])
}

func abs(a int) int {
	if a > 0{
		return a
	}
	return a * -1
}

func main() {
	res := allCellsDistOrder(2, 2, 0, 1)
	fmt.Print(res)
}
