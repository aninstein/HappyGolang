package main

import (
    "fmt"
    "strings"
)

const NMAX = 10

// dp代表的是字符串的个数
// 每一次放进去这个字符串就判断这个m和n还有没有值
// 状态转移方程二维的m、n数组，需要对mj和nj逆序
// dp[i][mj][nj] = max(dp[i-1][mj-num_0][nj-num_1]+, dp[i-1][mj][nj])
func findMaxForm(strs []string, m int, n int) int {
    if strs == nil {
        return 0
    }
    dataLen := len(strs)
    if dataLen == 0 {
        return 0
    }

    var dp [NMAX][NMAX]int
    for i := 0; i < dataLen; i++ {
        str := strs[i]
        num_0 := strNum(str, "0")
        num_1 := strNum(str, "1")
        j, k := m, n
        for j >= num_0 && k >= num_1 {
            dp[j][k] = max(dp[j-num_0][k-num_1]+1, dp[j][k])
            if j >= num_0 {
                j--
            }
            if k >= num_1 {
                k--
            }
        }

    }
    return dp[m][n]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func strNum(s, search string) int {
    return strings.Count(s, search)
}

func main() {
    var strArray []string
    strArray = append(strArray, "10", "1111", "111000001", "1", "0", "0111", "1011", "0")
    m, n := 5, 3
    fmt.Println(findMaxForm(strArray, m, n))
}
