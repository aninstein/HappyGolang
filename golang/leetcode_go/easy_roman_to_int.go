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
	for i := 1; i < dataLen; i++ {
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
//
//func main() {
//	romanToInt("111")
//}
