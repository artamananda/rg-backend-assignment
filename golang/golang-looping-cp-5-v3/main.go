package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseWord(str string) string {
	res := ""
	strSplit := strings.Split(str, " ")

	for i := 0; i < len(strSplit); i++{
		for j := len(strSplit[i]) - 1; j >= 0; j--{
			a := strSplit[i][0]
			if j == len(strSplit[i]) - 1{
				if unicode.IsUpper(rune(a)){
					res = res + strings.ToUpper(string(strSplit[i][j]))
				} else{
					res = res + strings.ToLower(string(strSplit[i][j]))
				}
			} else{
				res = res + strings.ToLower(string(strSplit[i][j]))
			}
		}
		if i < len(strSplit) - 1{
			res = res + " "
		}
	}


	return res // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseWord("Aku Sayang Ibu"))
	fmt.Println(ReverseWord("A bird fly to the Sky"))
}
