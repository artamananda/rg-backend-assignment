package main

import (
	"fmt"
)

func PhoneNumberChecker(number string, result *string) {
	if len(number) >= 10 && number[0] == '0' && number[1] == '8'{
		if number[2] == '1' && number[3] >= 49 && number[3] <= 53{
			*result = "Telkomsel"
		} else if number[2] == '1' && number[3] >= 54 && number[3] <= 57{
			*result = "Indosat"
		} else if number[2] == '2' && number[3] >= 49 && number[3] <= 51{
			*result = "XL"
		} else if number[2] == '2' && number[3] >= 55 && number[3] <= 57{
			*result = "Tri"
		} else if number[2] == '5' && number[3] >= 50 && number[3] <= 51{
			*result = "AS"
		} else if number[2] == '8' && number[3] >= 49 && number[3] <= 56{
			*result = "Smartfren"
		} else{
			*result = "invalid"
		}
	} else if len(number) >= 11 && number[0] == '6' && number[1] == '2' && number[2] == '8' {
		if number[3] == '1' && number[4] >= 49 && number[4] <= 53{
			*result = "Telkomsel"
		} else if number[3] == '1' && number[4] >= 54 && number[4] <= 57{
			*result = "Indosat"
		} else if number[3] == '2' && number[4] >= 49 && number[4] <= 51{
			*result = "XL"
		} else if number[3] == '2' && number[4] >= 55 && number[4] <= 57{
			*result = "Tri"
		} else if number[3] == '5' && number[4] >= 50 && number[4] <= 51{
			*result = "AS"
		} else if number[3] == '8' && number[4] >= 49 && number[4] <= 56{
			*result = "Smartfren"
		} else{
			*result = "invalid"
		}
	} else{
		*result = "invalid"
	}
}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "082011111111"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}
