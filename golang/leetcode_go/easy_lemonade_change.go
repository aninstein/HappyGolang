package main

import "fmt"

func lemonadeChange(bills []int) bool {
	dataLen := len(bills)
	if dataLen == 0 {
		return false
	} else if dataLen == 1 {
		return bills[0] == 5
	}
	if bills[0] != 5 {
		return false
	}
	surplus := 5
	for i:=1;i<dataLen;i++ {
		pay := bills[i]
		if pay != 5 {
			retPay := pay - 5
			if surplus < retPay {
				fmt.Println("pay", pay)
				fmt.Println("retPay", retPay)
				fmt.Println("surplus", surplus)
				return false
			}
			surplus += (surplus - retPay) + pay
		} else {
			surplus += 5
		}
	}
	return true
}

func main() {
	data := []int{5,5,10,10,20}
	fmt.Println(lemonadeChange(data))
}