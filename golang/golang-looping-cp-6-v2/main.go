package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	res := 0
	temp := 0
	var strNum = fmt.Sprintf("%d", numbers)
	
	for i := 0; i < len(strNum) - 1; i++{
		var j1,_ = strconv.Atoi(string(strNum[i]))
		var j2,_ = strconv.Atoi(string(strNum[i+1]))
		if temp < j1 + j2{
			temp = j1+j2
			var j3,_= strconv.Atoi(string(strNum[i]) + string(strNum[i+1]))
			res = j3
		}
	}
	
	return res
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(11223344))
}
