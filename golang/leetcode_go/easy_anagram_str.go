package main

import "fmt"


func isAnagram(s string, t string) bool {
	sLen := len(s)
	if sLen != len(t) {
		return false
	}

	sMap := make(map[int]int)
	tMap := make(map[int]int)
	for i:=97; i<=123; i++ {  // ascii
		sMap[i] = 0
		tMap[i] = 0
	}

	for i:=0; i<sLen; i++ {
		sMap[int(s[i])]++
		tMap[int(t[i])]++
	}

	for i:='a'; i<='z'; i++ {
		if sMap[int(i)] != tMap[int(i)] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println('z')

}
