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

// 对于上述的状态转移方程，其实我们可以做一些调整
// 我们设定tmp数组，对于tmp[i]，即当前最长递增子序列长度为i的，是到data数组中的哪一个数字
// 而tmp数组中存放的就是最长递增子序列
func tmpFunction(data []int) int {
    dataLen := len(data)
    if dataLen == 0 {
        return 0
    } else if dataLen == 1 {
        return 1
    }

    maxLen := 1
    var tmp [MAXN]int
    tmp[0] = data[0]
    for i:=1; i<dataLen+1; i++ {
        if data[i] > tmp[maxLen-1] {
            maxLen++
            tmp[maxLen-1] = data[i]
        } else {
            binSearch(tmp, data[i])
        }
    }
    return maxLen
}


// 二分查找并且替换data左边的元素
func binSearch(data [MAXN]int, index int) {
    dataLen := len(data)
    if dataLen ==0 || dataLen == 1 {
        return
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
            data[mid] = index
        }
    }
    data[left] = index
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


func main() {
    //data := createData(10)
    var data []int
    data = append(data, -2, -1)
    res := dpFunction(data)
    fmt.Println("dpFunction: ", res)

    //res = tmpFunction(data)
    //fmt.Println("tmpFunction: ", res)
}