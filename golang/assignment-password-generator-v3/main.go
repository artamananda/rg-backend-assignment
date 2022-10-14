package main

import(
	"fmt"
	"strings"
) 

func Reverse(str string) string {
	revStr := ""
	for i := len(str)-1; i >= 0; i--{
		revStr += string(str[i])
	}
	return revStr
}

func Generate(str string) string {
	str = Reverse(str)
	res := ""
	str = strings.ReplaceAll(str, " ", "")

	for i:=0; i < len(str); i++{
		if str[i] == 'A' || str[i] == 'I'|| str[i] == 'U' || str[i] == 'E' || str[i] == 'O'{
			if str[i] == 'A'{
				res += "e"
			} else if str[i] == 'E'{
				res += "i"
			} else if str[i] == 'I'{
				res += "o"
			} else if str[i] == 'O'{
				res += "u"
			} else{
				res += "a"
			}
		} else if str[i] == 'a' || str[i] == 'i'|| str[i] == 'u' || str[i] == 'e' || str[i] == 'o'{
			if str[i] == 'a'{
				res += "E"
			} else if str[i] == 'e'{
				res += "I"
			} else if str[i] == 'i'{
				res += "O"
			} else if str[i] == 'o'{
				res += "U"
			} else{
				res += "A"
			}
		} else if str[i] > 64 && str[i] < 91{
			res += strings.ToLower(string(str[i]))
		} else if str[i] > 96 && str[i] < 123{
			res += strings.ToUpper(string(str[i]))
		} else{
			res += string(str[i])
		}
		
	}

	return res
}

func CheckPassword(str string) string {
	check := ""
	lemah := false
	sedang := false
	count := 0

	for i := 0; i < len(str); i++{
		if (str[i] > 64 && str[i] < 91) || (str[i] > 47 && str[i] < 58) || (str[i] > 96 && str[i] < 123){
			lemah = true
			count++
		} else{
			lemah = false
			sedang = true
		}
	}

	if count == 0{
		sedang = false
		lemah = false
	}

	if len(str) >= 14 && sedang == true{
		check = "kuat"
	} else if len(str) >= 7 && sedang == true{
		check = "sedang"
	} else if len(str) >= 7 && lemah == true{
		check = "lemah"
	} else {
		check = "sangat lemah"
	}

	return check
}

func PasswordGenerator(base string) (string, string) {
	return Generate(base), CheckPassword(Generate(base))
}

func main() {
	data := "admin1238*(#@123i@DJHDJSUI@*(#)NDKJKJDSNJ" // bisa digunakan untuk melakukan debug

	res, check := PasswordGenerator(data)

	fmt.Println(res)
	fmt.Println(check)
}
