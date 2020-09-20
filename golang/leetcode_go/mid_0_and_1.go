package main

import (
    "fmt"
    "strings"
)

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

    var dp [dataLen+10][m+10][n+10]int
    for i := 1; i < dataLen+1; i++ {
        str := strs[i-1]
        num_0 := strNum(str, "0")
        num_1 := strNum(str, "1")
        for j:=0; j<=m; j++ {
            for k:=0; k <= n; k++ {
                if j >= num_0 && k >= num_1 {
                    dp[i][j][k] = max(dp[i-1][j-num_0][k-num_1]+1, dp[i][j][k])
                }
                dp[i][j][k] = max(dp[i][j][k], dp[i-1][j][k])
            }
        }
    }
    return dp[dataLen][m][n]
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
    strArray = append(strArray, "00", "000")
    m, n := 1, 10
    fmt.Println(findMaxForm(strArray, m, n))
}
