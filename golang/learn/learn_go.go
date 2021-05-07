package main

import "fmt"

func main() {
    var a [3]int
    a = [3]int{1, 2, 3}
    b := []int{1, 2, 3}
    fmt.Println(a, b)
    c := append(b, 1, 2, 3, 4, 5, 6, 7, 8, 9)  // 实际上并不是对b的append，而是生成了一个新的切片，切片b的长度还是不变
    fmt.Println(c, b)
    var s1 []int = make([]int, 10, 20)  // 使用内置函数生成切片，第一个参数了len为初始化切片的长度，第二个参数cap为此切片的最大长度
    fmt.Println(s1, len(s1), cap(s1))

    s2 := a[:]  // 这样截取一个数组也能够生成一个切片，其len和cap就是截取的数组长度
    fmt.Println(s2, len(s2), cap(s2))

    /*
    数组和切片的区别
    1. 声明
    数组声明：a := [10]int{···}  // 数组长度是固定的
    切片声明：b := []int{···}  // 切片的方括号是空的，长度根据初始化的时候赋值数据长度而定

    2. 赋值
    - 数组截取赋值给切片是值赋值
    - 切片直接赋值给切片是引用赋值
    */


    var countryCapitalMap map[string]string /*创建集合 */
    countryCapitalMap = make(map[string]string)

    /* map插入key - value对,各个国家对应的首都 */
    countryCapitalMap [ "France" ] = "巴黎"
    countryCapitalMap [ "Italy" ] = "罗马"
    countryCapitalMap [ "Japan" ] = "东京"
    countryCapitalMap [ "India" ] = "新德里"

    /*使用键输出地图值 */
    for country := range countryCapitalMap {
        fmt.Println(country, "首都是", countryCapitalMap [country])
    }

    /*查看元素在集合中是否存在 */
    capital, ok := countryCapitalMap["India"] /*如果确定是真实的,则存在,否则不存在 */
    fmt.Println(capital)
    fmt.Println(ok)
    if (ok) {
        fmt.Println("India 的首都是", capital)
    } else {
        fmt.Println("India 的首都不存在")
    }
}

/*
go的类
*/

/*
go的接口
// 定义接口
type interface_name interface {
    method_name1 [return_type]
    method_name2 [return_type]
    method_name3 [return_type]
    ...
    method_namen [return_type]
}

// 定义结构体
type struct_name struct {
    // variables
}

/* 实现接口方法
func (struct_name_variable struct_name) method_name1() [return_type] {
    // 方法实现
}
...
func (struct_name_variable struct_name) method_namen() [return_type] {
    // 方法实现
}

func main() {
    var phone Phone

    phone = new(NokiaPhone)
    phone.call()

    phone = new(IPhone)
    phone.call()

}

*/

type ListNode struct {
    index int
    Val int
    Next *ListNode
}

type ArrayList interface {
    insertLastNode(node *ListNode)  // 尾部插入节点
    insertNode(index int, node *ListNode)  // 从index插入节点
    deleteNode(index int)  // 删除index的节点
    getNode(index int) *ListNode  // 获取第index个节点
}

type ArrayListImpl struct {
    listLen int
    firstNode *ListNode
}

func (arrayListImpl ArrayListImpl) insertLastNode(node *ListNode) {
    /* 链表的尾插入法 */
    if arrayListImpl.firstNode == nil {
        arrayListImpl.firstNode = node
        arrayListImpl.listLen = 1
        return
    }
    count := 1
    var next *ListNode
    for true {
        next = arrayListImpl.firstNode.Next
        if next == nil{
            node.index = count
            arrayListImpl.firstNode.Next = node
            arrayListImpl.listLen = count + 1
            break
        }
        count++
    }
}

func (arrayListImpl ArrayListImpl) insertNode(index int, node *ListNode) {
    /* 链表的插入操作 */
    if index > arrayListImpl.listLen {
        return
    }
    if arrayListImpl.firstNode == nil {
        if index != 0 {
            return
        }
        arrayListImpl.firstNode = node
        return
    }
    if index == 0 {
        node.Next = arrayListImpl.firstNode
        node.index = 0
        arrayListImpl.firstNode = node
        arrayListImpl.listLen++
    } else {
        indexNode := arrayListImpl.getNode(index-1)
        node.Next = indexNode.Next
        node.index = indexNode.index
        indexNode.Next = node
        arrayListImpl.listLen++
    }
    upNode := arrayListImpl.getNode(index+1)
    for i:=index+1;i<arrayListImpl.listLen;i++ {
        if upNode != nil {
            upNode.index++
        }
        if upNode.Next != nil {
            upNode = upNode.Next
        }
    }
}

func (arrayListImpl ArrayListImpl) deleteNode(index int) {
    /* 链表的删除第index个节点 */
    if index > arrayListImpl.listLen {
        return
    }
    if arrayListImpl.firstNode == nil {
        return
    }
    if index == 0 {
        arrayListImpl.firstNode = arrayListImpl.firstNode.Next
    } else {
        indexNode := arrayListImpl.getNode(index-1)
        delNode := indexNode.Next
        indexNode.Next = delNode.Next
    }
    // 更新index，懒得写了
}

func (arrayListImpl ArrayListImpl) getNode(index int) *ListNode {
    /* 链表获取第index个配置 */
    if index > arrayListImpl.listLen {
        return nil
    }
    node := arrayListImpl.firstNode
    if node.index == index {
        return node
    }
    for true {
        node = node.Next
        if node == nil || node.index == index{
            return node
        }
    }
    return nil
}



