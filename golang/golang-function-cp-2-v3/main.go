package main

import (
	"fmt"
	"strings"
)

func CountVowelConsonant(str string) (int, int, bool) {
	str = strings.ToLower(str)
	vowel := 0
	consonant := 0
	boole := true

	for i := 0; i < len(str); i++{
		if str[i] == 'a' || str[i] == 'i' || str[i] == 'u' || str[i] == 'e' || str[i] == 'o'{
			vowel++
		} else if str[i] == ' ' || str[i] == ','{
			continue
		} else {
			consonant++
		}
	}

	if vowel == 0 || consonant == 0{
		boole = true
	} else{
		boole = false
	}

	return vowel, consonant, boole
}
// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("bbbbbbbb vvvvvvvv  ddddd sssss  wwww"))
}
