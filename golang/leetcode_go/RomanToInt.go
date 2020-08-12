package main
// 罗马数字转整数
// https://leetcode-cn.com/problems/roman-to-integer/
func romanToInt(s string) int {
    romanMap := make(map[string]int)
    romanMap["I"] = 1
    romanMap["V"] = 5
    romanMap["X"] = 10
    romanMap["L"] = 50
    romanMap["C"] = 100
    romanMap["D"] = 500
    romanMap["M"] = 1000
    dataLen := len(s)
    preNum := romanMap[string(s[0])]
    sum := preNum
    for i:=1; i < dataLen; i++ {
       char := string(s[i])
       num := romanMap[char]
       if num > preNum {
           // sum -= preNum
           // sum += num - preNum
           sum += num - (preNum * 2)
       } else {
           sum += num
           preNum = num
       }
    }
    return sum
}

func main()  {
    romanToInt("111")
}


/*
package java.java;

public class AlgorithmSort {

    public static void main(String[] args) {
        int[] data = {1, 2, 7, 1, -9, 100, 5400, -471, -1, 0, 0, 21, 3};
        bubbleSort(data);
        for (int i : data) {
            System.out.print(i + " ");
        }


    }

    private static int[] bubbleSort(int[] data){

        if(data == null || data.length == 0){
            return data;
        }

        int dataLen = data.length;
        boolean flag = true;
        for (int i = 0; i < dataLen; i++) {
            for (int j = i+1; j < dataLen; j++) {
                if(data[i] > data[j]){
                    int tmp = data[i];
                    data[i] = data[j];
                    data[j] = tmp;
                    flag = false;
                }
            }
            if(flag) break;
        }
        return data;
    }

}
*/ 