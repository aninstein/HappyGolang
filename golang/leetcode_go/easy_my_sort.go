package main

import (
	"fmt"
	"sort"
)

/*
题目：距离顺序排列矩阵单元格
题目链接：https://leetcode-cn.com/problems/matrix-cells-in-distance-order/
[[(0, 0), (0, 1), (0, 2), (0, 3), (0, 4), (0, 5)]
[(1, 0), (1, 1), (1, 2), (1, 3), (1, 4), (1, 5)]
[(2, 0), (2, 1), (2, 2), (2, 3), (2, 4), (2, 5)]
[(3, 0), (3, 1), (3, 2), (3, 3), (3, 4), (3, 5)]
[(4, 0), (4, 1), (4, 2), (4, 3), (4, 4), (4, 5)]
[(5, 0), (5, 1), (5, 2), (5, 3), (5, 4), (5, 5)]]
*/
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	dataMap := make(map[float64][]int)
	index := []int{r0, c0}
	salt := 0.000001
	offset := 0.0
	for i:=0; i<R; i++ {
		for j:=0; j<C; j++ {
			cell := []int{i, j}
			dis := manhatan(index, cell)
			_, ok := dataMap[dis]
			if ok {
				offset += salt
				dataMap[dis+offset] = cell
			} else {
				dataMap[dis] = cell
			}
		}
	}

	var keyList []float64
	for k, _ := range dataMap {
		keyList = append(keyList, k)
	}

	sort.Float64s(keyList)
	var mt [][]int
	for i:=0; i<len(keyList); i++ {
		mt = append(mt, dataMap[keyList[i]])
	}
	return mt
}

func manhatan(a, b []int) float64 {
	return float64(abs(a[0] - b[0]) + abs(a[1] - b[1]))
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return a * -1
}

func allCellsDistOrder2(R int, C int, r0 int, c0 int) [][]int {
	layer := 1
	var mt [][]int
	mt = append(mt, []int{r0, c0})
	for layer < R || layer < C {
		step := 0
		up := r0 - layer
		down := r0 + layer
		left := c0 - layer
		right := c0 + layer
		for step < layer {
			// 上
			if up >= 0 {
				r := c0 + step
				if r < right && r < C {
					mt = append(mt, []int{up, r})
				}
				l := c0 - step
				if l > left && l >= 0 && step > 0{
					mt = append(mt, []int{up, l})
				}
			}

			// 下
			if down < R {
				r := c0 + step
				if r <= right && r < C {
					mt = append(mt, []int{down, r})
				}
				l := c0 - step
				if l >= left && l >= 0 && step > 0{
					mt = append(mt, []int{down, l})
				}
			}

			// 左，不处理角点
			if left >= 0 {
				d := r0 + step
				if d < down && d < R {
					mt = append(mt, []int{d, left})
				}
				u := r0 - step
				if u > up && u >= 0 && step > 0{
					mt = append(mt, []int{u, left})
				}
			}

			// 右，不处理角点
			if right < C {
				d := r0 + step
				if d < down && d < R {
					mt = append(mt, []int{d, right})
				}
				u := r0 - step
				if u > up && u >= 0 && step > 0 {
					mt = append(mt, []int{u, right})
				}
			}
			step++
		}
		if up >=0 && left >= 0{
			mt = append(mt, []int{up, left})
		}
		if up >=0 && right < C{
			mt = append(mt, []int{up, right})
		}
		if down < R && left >= 0{
			mt = append(mt, []int{down, left})
		}
		if down < R && right < C{
			mt = append(mt, []int{down, right})
		}
		layer++
	}
	return mt
}

func main() {
	R, C := 1, 2
	r0, c0 := 0, 0
	res1 := allCellsDistOrder(R, C, r0, c0)
	fmt.Print(len(res1))
	fmt.Println(res1)
	res2 := allCellsDistOrder2(R, C, r0, c0)
	fmt.Print(len(res2))
	fmt.Println(res2)
}