package main

import (
	"fmt"
)

/*
题目：283. 移动零
题目链接：https://leetcode-cn.com/problems/move-zeroes/

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
*/

// 自己写的
func moveZeroes(nums []int) []int {
	numsLen := len(nums)
	var index int
	for i:=numsLen-1; i>=0; i-- {
		if nums[i] != 0 {
			index = i
			break
		}
	}

	for i:=index-1; i>=0; i-- {
		if nums[i] == 0{
			for j:=i; j<index; j++ {
				nums[j] = nums[j+1]
			}
			nums[index] = 0
			index--
		}
	}
	return nums
}

// 官方题解
func moveZeroes2(nums []int) []int {
	numsLen, right, left := len(nums), 0, 0
	for right < numsLen {
		if nums[right] != 0 {
			nums[right], nums[left] = nums[left], nums[right]
			left++
		}
		right++
	}
	return nums
}

func main() {
	ll := []int{0,0,1}
	res := moveZeroes(ll)
	fmt.Println(res)
}