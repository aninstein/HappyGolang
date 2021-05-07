package main

import (
	"bytes"
	"fmt"
	"strconv"
)

/*
* 题目：汇总区间
* 题目链接：https://leetcode-cn.com/problems/summary-ranges/
*/


func summaryRanges(nums []int) []string {
	dataLen := len(nums)
	if dataLen == 0 {
		return []string{}
	}

	next_num := nums[0] + 1
	var retData []string
	for i:=0; i < dataLen; i++ {
		j := i + 1
		if j > dataLen - 1 {
			retData = append(retData, strconv.FormatInt(int64(nums[dataLen - 1]), 10))
			return retData
		}
		for next_num == nums[j] {
			next_num++
			j++
			if j > dataLen - 1 {
				retData = append(retData, getRangeStr(nums[i], nums[j-1]))
				return retData
			}
		}
		retData = append(retData, getRangeStr(nums[i], nums[j-1]))
		next_num = nums[j] + 1
		i = j - 1
	}
	return retData
}

func getRangeStr(left, right int) string {
	if left == right {
		return strconv.FormatInt(int64(left), 10)
	}
	var bt bytes.Buffer
	// 向bt中写入字符串
	bt.WriteString(strconv.FormatInt(int64(left), 10))
	bt.WriteString("->")
	bt.WriteString(strconv.FormatInt(int64(right), 10))
	//获得拼接后的字符串
	return bt.String()
}

func main() {
	var data []int
	data = append(data, 0,2,3,4,6,8,9)
	res := summaryRanges(data)
	fmt.Println(res)
}