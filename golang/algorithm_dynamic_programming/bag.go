package main

import (
    "fmt"
    "math/rand"
    "time"
)

const ROW = 10000
const COL = 10000

func createBagData(number int, bagType string) map[string][]int {
    rand.Seed(time.Now().UnixNano()) // 时间做随机种子
    dataMap := make(map[string][]int)
    dataMap["weight"] = []int{}
    dataMap["value"] = []int{}
    dataMap["number"] = []int{}
    for i := 0; i < number; i++ {
        if bagType == "multiple" {
            appendData(dataMap, rand.Intn(9)+1, rand.Intn(99)+1, rand.Intn(19)+1)
        } else if bagType == "complete" {
            appendData(dataMap, rand.Intn(9)+1, rand.Intn(99)+1, 9999)
        } else {
            appendData(dataMap, rand.Intn(9)+1, rand.Intn(99)+1, 1)
        }
    }
    return dataMap
}

func appendData(data map[string][]int, weight, value, number int) {
    data["weight"] = append(data["weight"], weight)
    data["value"] = append(data["value"], value)
    data["number"] = append(data["number"], number)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

//    01背包问题
//    问题描述：有n件物品和容量为m的背包 给出i件物品的重量以及价值 求解让装入背包的物品重量不超过背包容量 且价值最大 。
//    特点:这是最简单的背包问题，特点是每个物品只有一件供你选择放还是不放。
func pak0And1(data map[string][]int, totalWeight int) int {
    weightList := data["weight"]
    valueList := data["value"]
    numberList := data["weight"]
    totalNumber := len(weightList)
    fmt.Println("Two dim function result: ")
    res := pak0And1TwoDim(totalNumber, totalWeight, valueList, weightList, numberList)
    fmt.Println(res)
    fmt.Println("One dim function result: ")
    res = pak0And1OneDim(totalNumber, totalWeight, valueList, weightList, numberList)
    fmt.Println(res)
    return res
}

// 《二维实现方法》，状态转移方程：
// dp[i][j] = max(dp[i-1][j-w[i]]+v[i], dp[i-1][j])
func pak0And1TwoDim(totalNum, totalWeight int, v, w, num []int) int {
    // 之所以加1是因为数据从1开始计算，为什么从1开始计算呢，是因为个数和重量都是表示实际意义的
    row := totalNum + 1
    col := totalWeight + 1
    var dp = [ROW][COL]int{}
    for i := 1; i < row; i++ {
        if i == totalNum {
            break
        }
        for j := 1; j < col; j++ {
            if j >= w[i] {
                dp[i][j] = max(dp[i-1][j-w[i]]+v[i], dp[i-1][j])
            } else {
                dp[i][j] = dp[i-1][j]
            }
        }
    }
    return dp[totalNum-1][totalWeight]
}

// 《一维实现方法》，状态转移方程
// dp[j] = max(dp[j-w[i]]+v[i], dp[j]) | j逆序
func pak0And1OneDim(totalNum, totalWeight int, v, w, num []int) int {
    // 之所以加1是因为数据从1开始计算，为什么从1开始计算呢，是因为个数和重量都是表示实际意义的
    row := totalNum + 1
    col := totalWeight + 1
    var dp = make([]int, col)
    for i := 1; i < row; i++ {
        if i == totalNum {
            break
        }
        // 为什么倒序？是因为这个地方的下标"j-w[i]"是减法
        // 会回头找第i轮已经被更新的数据进行比较，实际上我们应该找的是第i-1轮的旧数据进行比较
        for j := col - 1; j >= w[i]; j-- {
            dp[j] = max(dp[j-w[i]]+v[i], dp[j])
        }
    }
    return dp[totalWeight]
}

//    完全背包问题
//    问题描述：有n件物品和容量为m的背包 给出i件物品的重量以及价值 求解让装入背包的物品重量不超过背包容量 且价值最大 。
//    特点：题干看似与01一样 但它的特点是每个物品可以无限选用。
func pakComplete(data map[string][]int, totalWeight int) int {
    weightList := data["weight"]
    valueList := data["value"]
    numberList := data["weight"]
    totalNumber := len(weightList)
    fmt.Println("Two dim and K function result: ")
    res := pakCompleteKTwoDim(totalNumber, totalWeight, valueList, weightList, numberList)
    fmt.Println(res)

    fmt.Println("Two dim not K function result: ")
    res = pakCompleteTwoDim(totalNumber, totalWeight, valueList, weightList, numberList)
    fmt.Println(res)

    fmt.Println("One dim function result: ")
    res = pakCompleteOneDim(totalNumber, totalWeight, valueList, weightList, numberList)
    fmt.Println(res)

    fmt.Println("One dim function pakCompleteOneDimDataOptimizaion result: ")
    res = pakCompleteOneDimDataOptimizaion(totalNumber, totalWeight, valueList, weightList, numberList)
    fmt.Println(res)
    return res
}

// 《二维带K循环解法》
// 对每个物品：
// 选择放进去：可以放进去1, 2, 3 ··· k, k-1个，不能是无限个，因为背包空间有限
// 选择不放进去：即放进去0个
// 状态转移方程：
// dp[i][j] = max(dp[i-1][j-k*w[i]] + k*v[i], dp[i-1][j])
func pakCompleteKTwoDim(totalNum, totalWeight int, v, w, num []int) int {
    row := totalNum + 1
    col := totalWeight + 1
    var dp = [ROW][COL]int{}
    for i := 1; i < row; i++ {
        if i == totalNum {
            break
        }
        for j := w[i]; j < col; j++ {
            dp[i][j] = dp[i-1][j]
            for k := 1; k*w[i] <= j; k++ {
                dp[i][j] = max(dp[i-1][j-k*w[i]]+k*v[i], dp[i-1][j])
            }
        }
    }
    return dp[totalNum-1][totalWeight]
}

// 《 二维不带K循环解法（时间复杂度优化）》
// 数学方法对k进行转化后发现，其实这个过程不过是第i轮与第i轮数据重复比较的过程
// 具体逻辑可以查看链接：https://blog.csdn.net/aninstein/article/details/108091495
// 最终状态转移方程
// dp[i][j] = max(dp[i][j-w[i]] + v[i], dp[i-1][j])
func pakCompleteTwoDim(totalNum, totalWeight int, v, w, num []int) int {
    row := totalNum + 1
    col := totalWeight + 1
    var dp = [ROW][COL]int{}
    for i := 1; i < row; i++ {
        if i == totalNum {
            break
        }
        for j := w[i]; j < col; j++ {
            dp[i][j] = max(dp[i][j-w[i]]+v[i], dp[i-1][j])
        }
    }
    return dp[totalNum-1][totalWeight]
}

// 《一维解法（空间复杂度优化）》
// 从上面可以了解到，实际上完全背包算法就是第i轮与第i轮数据重复比较的过程，上面01背包的优化方法
// 当j逆序遍历的时候，可以保证到的是每次进行比较的是第i轮与第i-1轮数据比较的过程
// 当j正序遍历的时候，则可以保证到的是每次进行比较的是第i轮与第i轮数据比较的过程
// 因此状态转移方程可以改为：
// dp[j] = max(dp[j-w[i]+v[i]], dp[j]) | j正序
func pakCompleteOneDim(totalNum, totalWeight int, v, w, num []int) int {
    row := totalNum + 1
    col := totalWeight + 1
    var dp = [COL]int{}
    for i := 1; i < row; i++ {
        if i == totalNum {
            break
        }
        for j := w[i]; j < col; j++ {
            dp[j] = max(dp[j-w[i]]+v[i], dp[j])
        }
    }
    return dp[totalWeight]
}

// 《一维解法（空间复杂度优化）》
// 上面只是对于算法的优化，实际上我们还可以通过对数据源的优化
// 完全背包算法可以去除掉同重量的，只取价值高的物品，这是由于物品可以无线选取
// 但是01背包不可以，因为有可能两个物品都会需要
func pakCompleteOneDimDataOptimizaion(totalNum, totalWeight int, v, w, num []int) int {
    vwMap := make(map[int] int)
    for i:=0; i<len(v); i++ {
        val, ok := vwMap[w[i]]
        if (ok && val < v[i]) || !ok {
            vwMap[w[i]] = v[i]
        }
    }

    var tv, tw []int
    for key, val := range vwMap {
        tv = append(tv, val)
        tw = append(tw, key)
    }

    row := len(tv) + 1
    col := totalWeight + 1
    var dp = [COL]int{}
    for i := 1; i < row; i++ {
        if i == totalNum {
            break
        }
        for j := w[i]; j < col; j++ {
            dp[j] = max(dp[j-w[i]]+v[i], dp[j])
        }
    }
    return dp[totalWeight]
}


//    多重背包问题
//    问题描述：有n件物品和容量为m的背包 给出i件物品的重量以及价值 还有数量 求解让装入背包的物品重量不超过背包容量 且价值最大 。
//    特点 ：它与完全背包有类似点 特点是每个物品都有了一定的数量。
func pakMultiple(data map[string][]int, totalWeight int) int {
    weightList := data["weight"]
    valueList := data["value"]
    numberList := data["weight"]
    totalNumber := len(weightList)
    fmt.Println("Two dim and K function result: ")
    res := pakMultipleKTwoDim(totalNumber, totalWeight, valueList, weightList, numberList)
    fmt.Println(res)

    fmt.Println("One dim and K function result: ")
    res = pakMultipleKOneDim(totalNumber, totalWeight, valueList, weightList, numberList)
    fmt.Println(res)

    fmt.Println("One pakMultipleBin function result: ")
    res = pakMultipleBin(totalNumber, totalWeight, valueList, weightList, numberList)
    fmt.Println(res)

    return res
}

// 《二维带K循环解法》
// 对每个物品：
// 选择放进去：可以放进去1, 2, 3 ··· k, k-1个，k范围：0 < k && k*w[i] <= j && k <= num[i]
// 选择不放进去：即放进去0个
// 状态转移方程：
// dp[i][j] = max(dp[i-1][j-k*w[i]] + k*v[i], dp[i-1][j]) | 0 < k && k*w[i] <= j && k <= num[i]
func pakMultipleKTwoDim(totalNum, totalWeight int, v, w, num []int) int {
    row := totalNum + 1
    col := totalWeight + 1
    var dp = [ROW][COL]int{}
    for i := 1; i < row; i++ {
        if i == totalNum {
            break
        }
        for j := 0; j < col; j++ {
            dp[i][j] = dp[i-1][j]
            for k := 0; k*w[i] <= j && k <= num[i]; k++ {
                dp[i][j] = max(dp[i][j-k*w[i]]+k*v[i], dp[i][j])
            }
        }
    }
    return dp[totalNum-1][totalWeight]
}

// 《一维带K循环解法·空间复杂度优化》
// 对每个物品：
// 选择放进去：可以放进去1, 2, 3 ··· k, k-1个，k范围：0 < k && k*w[i] <= j && k <= num[i]
// 选择不放进去：即放进去0个
// 由于每一轮K的范围都不相同，因此不能够当做公因式进行提取，故不能去除掉k
// 当j逆序遍历的时候，可以保证到的是每次进行比较的是第i轮与第i-1轮数据比较的过程
// 当j正序遍历的时候，则可以保证到的是每次进行比较的是第i轮与第i轮数据比较的过程
// 因此此处需要的是逆序遍历，状态转移方程：
// dp[j] = max(dp[j-k*w[i]+k*v[i]], dp[j]) | j逆序
func pakMultipleKOneDim(totalNum, totalWeight int, v, w, num []int) int {
    row := totalNum + 1
    col := totalWeight + 1
    var dp = [ROW]int{}
    for i := 1; i < row; i++ {
        if i == totalNum {
            break
        }
        for j := col; j >= 0; j-- {
            for k := 0; k*w[i] <= j && k <= num[i]; k++ {
                dp[j] = max(dp[j-k*w[i]]+k*v[i], dp[j])
            }
        }
    }
    return dp[totalWeight]
}

// 《把多重背包转换成01背包，二进制优化法》
// 对元数据本身进行操作
// 我们把k分解成，【二的次方的和】+【剩余常数】
// 改装后的元数据，其实就只剩下
func pakMultipleBin(totalNum, totalWeight int, v, w, num []int) int {
    count := 1  // 表示数组的长度
    var tv, tw [COL]int
    for i:=0; i<totalNum; i++ {
        a, b, s := v[i], w[i], num[i]
        // 把当前的s进行分解，看看这个取值k可以到多少范围
        k := 1
        for k <= s {
            tv[count] = a * k
            tw[count] = b * k
            s -= k
            k <<= 1
            count++
        }

        // 这个s没完全被分解完的部分，即上述的常数c
        if s != 0 {
            tv[count] = a * k
            tw[count] = b * k
            count++
        }
    }

    // 至此，我们的数据已经归整完毕
    // tv，tw：存放的是，当2次幂个物品被放进来的时候，的价值，下标为count，这时候只需要进行01背包运算即可
    var dp [COL]int
    for i:=1; i<=count; i++ {
        for j:=totalWeight; j>=tw[i]; j-- {
            input := dp[j-tw[i]]+tv[i]
            notput := dp[j]
            dp[j] = max(input, notput)
        }
    }
    return dp[totalWeight]
}


func main() {
  fmt.Println("<<<<< For 0 and 1 pak >>>>>")
  data := createBagData(10, "01")
  fmt.Println(data)
  pak0And1(data, 30)

  fmt.Println()
  fmt.Println("<<<<< For complete pak >>>>>")
  data = createBagData(10, "complete")
  fmt.Println(data)
  pakComplete(data, 30)

  fmt.Println()
  fmt.Println("<<<<< For multiple pak >>>>>")
  data = createBagData(10, "multiple")
  fmt.Println(data)
  pakMultiple(data, 30)
}


//const N = 11010
//const M = 2010
//
//func main()  {
//    var n, m int
//    var v, w [N]int
//    fmt.Scanln(&n, &m)
//    cnt := 1
//    for i:=0; i<=n; i++ {
//        var a, b, s int
//        fmt.Scanln(&a, &b, &s)
//        k := 1
//        for k <= s {
//            v[cnt] = a*k
//            w[cnt] = b*k
//            s -= k
//            k <<= 1
//            cnt++
//        }
//
//        if s != 0 {
//            v[cnt] = a*s
//            w[cnt] = b*s
//        }
//    }
//    var dp [M]int
//    for i:=1; i<=cnt; i++ {
//        for j:=m; j>=v[i]; j-- {
//            dp[j]=max(dp[j],dp[j-v[i]]+w[i])
//        }
//    }
//    fmt.Println(dp[m])
//}
