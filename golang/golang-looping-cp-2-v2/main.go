package main

import "fmt"

// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {
	var res string = ""
	for i:=len(str) - 1; i >= 0; i--{
		res = res + string(str[i])
		if i > 0 && str[i] != ' ' && str[i-1] != ' '{
			res = res + "_"
		}
	}
	return res // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseString("Hello World"))
}
