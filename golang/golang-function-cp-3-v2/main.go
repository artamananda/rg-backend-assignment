package main

import (
	"fmt"
	"strings"
	"sort"
)

func FindShortestName(names string) string {
	res := ""

	for i := 0; i < len(names); i++{
		if names[i] == ' '{
			arr := strings.Split(names, " ")
			sort.Slice(arr, func(i, j int) bool {
				l1, l2 := len(arr[i]), len(arr[j])
				if l1 != l2 {
					return l1 < l2
				}
				return arr[i] < arr[j]
			})
			res = arr[0]
			break
		} else if names[i] == ','{
			arr := strings.Split(names, ",")
			sort.Slice(arr, func(i, j int) bool {
				l1, l2 := len(arr[i]), len(arr[j])
				if l1 != l2 {
					return l1 < l2
				}
				return arr[i] < arr[j]
			})
			res = arr[0]
			break
		} else if names[i] == ';'{
			arr := strings.Split(names, ";")
			sort.Slice(arr, func(i, j int) bool {
				l1, l2 := len(arr[i]), len(arr[j])
				if l1 != l2 {
					return l1 < l2
				}
				return arr[i] < arr[j]
			})
			res = arr[0]
			break
		} else{
			continue
		}
	}

	

	return res // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Ari,Aru,Ara,Andi,Asik"))
}
