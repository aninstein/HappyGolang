package main

import (
    "fmt"
    "math/rand"
    "sort"
    "time"
)

const DATALEN  = 20000

var TIMELIST map[int] string


func createData() []int {
    var retData [DATALEN]int
    for i:=0; i<DATALEN; i++ {
        retData[i] = rand.Intn(99999) + 1
    }
    return retData[:]
}


func bubbleSort(data []int) []int{
    if data == nil || len(data) == 0 {
        return []int{}
    }
    dataLen := len(data)
    for i := 0; i < dataLen; i++{
        for j := i + 1; j < dataLen; j++{
            if data[i] > data[j] {
                data[i], data[j] = data[j], data[i]
            }
        }
    }
    return data
}


func selectSort(data []int) []int{
    if data == nil || len(data) == 0{
        return []int{}
    }

    dataLen := len(data)
    for i := 0; i < dataLen; i++ {
        tmp := i
        for j := i+1; j < dataLen; j++ {
            if data[tmp] > data[j] {
                tmp = j
            }
        }
        if tmp != i {
            data[i], data[tmp] = data[tmp], data[i]
        }
    }
    return data
}


func insertSort(data []int) []int {
    if data == nil || len(data) == 0 {
        return []int{}
    }

    dataLen := len(data)
    for i := 1; i < dataLen; i++ {
        if data[i-1] > data[i] {
            tmp := data[i]
            for j := i; j >= 0; j-- {
                if j == 0 {
                    data[j] = tmp
                    break
                }

                if data[j-1] > tmp {
                    data[j] = data[j-1]  //挪动字段
                } else {
                    data[j] = tmp
                    break
                }
            }
        }
    }
    return data
}


// 希尔排序，插入排序改进版本
//    1. 把队列按照步长，分序列先做直接插入
//    2. 缩小等分步长做直接插入
//    3. 等分步长为1时候停止
func shellSort(data []int) []int {
    if data == nil || len(data) == 0 {
        return []int{}
    }

    dataLen := len(data)
    step := int((dataLen - 1) / 2)
    for step >= 1{
        step = int(step)
        for i := step; i < dataLen; i++ {
            tmp := data[i]
            j := i - step
            for j >= 0 && data[j] > tmp{  // 就是插入排序的挪动元素的步骤
                data[j+step] = data[j]
                j -= step
            }
            data[j+step] = tmp  // 完成元素的交换
        }
        step /= 2
    }
    return data
}

// 快速排序1
// 这种方法虽然便于理解，但是耗费空间比较多
func fastSort1(data []int) []int {
    if data == nil || len(data) == 0 {
        return []int{}
    }

    dataLen := len(data)
    var left, mid, right []int
    index := data[0]
    for i := 0; i < dataLen; i++ {
        if data[i] > index{
            right = append(right, data[i])
        }else if data[i] < index {
            left = append(left, data[i])
        }else {
            mid = append(mid, data[i])
        }
    }
    leftList := append(fastSort1(left), mid...)
    return append(leftList, fastSort1(right)...)
}


// 快速排序2
// 类似java和c那种，数组长度不能改变，没有类似python的list或者go的切片那样的数据类型
// 这种情况下就是充分的使用数组之间数据交换的方式进行数据排序
// 这种方法主要是不断的寻找最佳标定位置的方法
func fastSort2(data []int, left, right int) {
    if left < right{
        pos := position(data, left, right)
        fastSort2(data, left,  pos-1)
        fastSort2(data, pos+1,  right)
    }
}

func position(data []int, left, right int) int {
    index := data[right]  // 选用最后一个值作为主元
    i := left - 1  // 初始化标志位为第一个，之所以要减1，是因为传入参数是+1的
    for j := left; j < right; j++ {
        if index > data[j] {
            i++  // 标志位往前移动
            swap(data, i, j)  // 因为data[j]是小于index的，所以放到左边，且标志位往前移动
        }
    }
    swap(data, i+1, right)  // 求得到当前的标志位与主元交换完成
    return i+1
}

func mergeSort(data []int) []int {
    if data == nil {
        return []int{}
    }
    dataLen := len(data)
    if dataLen == 0{
        return []int{}
    } else if dataLen == 1{
        return data
    }


    haft := int(dataLen / 2)
    left := mergeSort(data[:haft])
    right := mergeSort(data[haft:])
    return merge(left, right)
}

func merge(left, right []int) []int {
    leftLen := len(left)
    rightLen := len(right)
    var newList []int
    i := 0
    j := 0
    for i < leftLen && j < rightLen {
        if left[i] < right[j] {
            newList = append(newList, left[i])
            i++
        }else {
            newList = append(newList, right[j])
            j++
        }
    }
    newList = append(newList, left[i:]...)
    newList = append(newList, right[j:]...)
    return newList
}


// 公用函数
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}


func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


func swap(data []int, i, j int) {
    data[i], data[j] = data[j], data[i]
}

func checkResult(data, right []int) bool {
    dataLen := len(data)
    rightLen := len(right)

    if dataLen != rightLen {
        return false
    }

    for i:=0; i<dataLen; i++ {
        if data[i] != right[i] {
            return false
        }
    }
    return true
}


func sortTestFunc(sortFunc func(data []int) []int, data, rightResult []int, name string)  {
    var usData []int
    usData = append(usData, data...)
    fmt.Println("# "+ name +" start >>>>>>>>>>>>>> ")
    start := time.Now().UnixNano()
    res := sortFunc(usData)
    total := time.Now().UnixNano() - start
    TIMELIST[int(total)] = name
    fmt.Println("# "+ name +" end, total time: ", total)
    fmt.Println("# "+ name + "result: ", res)
    fmt.Println("# "+ name +" result is: ", checkResult(res, rightResult))
    fmt.Println()
}


func main() {
    data := createData()
    var right []int
    right = append(right, data...)
    sort.Ints(right)
    fmt.Println("data: ", data)
    fmt.Println("right: ", right)
    TIMELIST = make(map[int] string)
    sortTestFunc(bubbleSort, data, right, "bubbleSort")
    sortTestFunc(selectSort, data, right, "selectSort")
    sortTestFunc(insertSort, data, right, "insertSort")
    sortTestFunc(shellSort, data, right, "shellSort")
    sortTestFunc(fastSort1, data, right, "fastSort1")
    sortTestFunc(mergeSort, data, right, "mergeSort")

    var keyList []int
    for key := range TIMELIST {
        keyList = append(keyList, key)
    }
    sort.Ints(keyList)


    for i:=0; i<len(keyList); i++ {
        key := keyList[i]
        fmt.Println("time: ", key, "func name: ", TIMELIST[key])
    }
}