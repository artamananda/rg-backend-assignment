package main

import (
	"fmt"
	"strings"
)

func ChangeOutput(data []string) map[string][]string {
	res := map[string][]string{}
	i := 0
	for i < len(data) {
		arr := strings.Split(data[i], "-")
		arr2 := strings.Split(data[i+1], "-")
		str := arr[3] + " " + arr2[3]
		if arr[0] == arr2[0] && arr[1] == arr2[1] {
			res[arr[0]] = append(res[arr[0]], str)
		} else {
			res[arr[0]] = append(res[arr[0]], arr[3])
			res[arr2[0]] = append(res[arr2[0]], arr2[3])
		}

		i += 2
	}

	return res // TODO: replace this
}

func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar", "phone-0-first-081234567890", "phone-1-first-081234567891"}
	fmt.Println(ChangeOutput(data))
}
