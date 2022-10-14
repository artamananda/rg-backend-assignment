package main

import (
	"fmt"
)

func DateFormat(day, month, year int) string {
	var months = [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	var res string = ""
	if day < 10{
		res += "0" + fmt.Sprintf("%d", day) + "-" + months[month-1] + "-" + fmt.Sprintf("%d", year)
	} else{
		res += fmt.Sprintf("%d", day) + "-" + months[month-1] + "-" + fmt.Sprintf("%d", year)
	}

	return res // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
}
