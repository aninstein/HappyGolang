package main

import "fmt"

/*
题目：加油站
题目链接：https://leetcode-cn.com/problems/gas-station/

输入:
gas  = [1,2,3,4,5]
cost = [3,4,5,1,2]

输出: 3

解释:
从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
因此，3 可为起始索引。
*/

func canCompleteCircuit(gas []int, cost []int) int {
	roadLen := len(gas)
	for i:=0; i<roadLen; i++ {
		sum := gas[i]
		preCost := cost[i]
		if sum < preCost {
			continue
		}
		for j:=1; j< roadLen + 2; j++ {
			index := getIndex(i, j, roadLen)
			sum -= preCost
			sum += gas[index]
			now := cost[index]
			if sum < now {
				break
			}
			preCost = now
			if index == i {
				if j != 1 {
					return i
				} else {
					return -1
				}
			}
		}
	}
	return -1
}

func getIndex(i, j, roadLen int) int {
	index := i + j
	if index < roadLen {
		return index
	}
	return index - roadLen
}

func main() {
	gas  := []int{1,2,3,4,5}
	cost := []int{3,4,5,1,2}
	res := canCompleteCircuit(gas, cost)
	fmt.Print(res)
}
