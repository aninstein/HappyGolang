package main
/*
* 题目：最长上升字序列
* 题目链接：https://leetcode-cn.com/problems/longest-increasing-subsequence/
*
*
*
*/

const NMAX = 10005

func lengthOfLIS(nums []int) int {
	dataLen := len(nums)
	if dataLen == 0 {
		return 0
	} else if dataLen == 1 {
		return 1
	} else if dataLen == 2 {
		if nums[0] < nums[1] {
			return 2
		} else {
			return 1
		}
	}
	// 状态转移方程
	// dp[i] = max(dp[j])+1, 其中0≤j<i且num[j]<num[i]
	// dp表示在这个数列以第i个数字结尾的递增字序列长度
	// 由于dp[j]表示的是以num[j]结尾的递增字序列长度，那么当num[i]>num[j]的时候，就可以进行状态转移，长度+1
	var dp [NMAX]int
	maxVal := 0
	dp[0] = 1
	for i := 1; i < dataLen; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
		maxVal = max(maxVal, dp[i])
	}
	return maxVal
}
