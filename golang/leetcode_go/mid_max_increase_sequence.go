package main

import (
    "fmt"
    "math/rand"
)

const MAXN = 100001

func createData(number int) []int {
    var data []int
    for i:=0; i<number; i++ {
        data = append(data, rand.Intn(99) + 1)
    }
    fmt.Println(data)
    return data
}


// 状态转移方程
// dp[i] = MAX(dp[j]+1,dp[i]), 初始化dp[i]=1
// 因为如果某个值比前面的值小，则它永远不会被这一轮的选上，即
// if i < j && arr[i] > arr[j], 则对于i这一轮，j永远不会被选上
// 此时dp[i] = max(dp[i]) + 1
func dpFunction(data []int) int {
    if data == nil{
        return 0
    }
    dataLen := len(data)
    if dataLen == 0 {
        return 0
    }
    var dp [MAXN]int
    maxLen := 0
    for i:=0; i<dataLen; i++ {
        dp[i] = 1
        maxNum := data[i]
        for j:=0; j<i; j++ {
            if maxNum < data[j] {
                dp[i] = max(dp[j]+1, dp[i])
            }
        }
        maxLen = max(maxLen, dp[i])
    }
    return maxLen
}

func binSearch(data []int, index int) int {
    dataLen := len(data)
    if dataLen ==0 || dataLen == 1 {
        return 0
    }
    left := 0
    right := dataLen - 1
    var mid int
    for left <= right {
        mid = (left + right) >> 1  // 左移1位即除以2
        if data[mid] > index {
            left = mid + 1
        } else if data[mid] < index {
            right = mid - 1
        } else {
            return mid
        }
    }
    return mid
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


func main() {
    data := createData(10)
    res := dpFunction(data)
    fmt.Println(res)
}