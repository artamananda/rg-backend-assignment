package main

import "fmt"
import "strings"

func CountingLetter(text string) int {
	// unreadable letters = R, S, T, Z
	var res int = 0
	var t string = strings.ToUpper(text)
	
	for i := 0; i < len(t); i++{
		if t[i] == 'R'{
			res++
		} else if t[i] == 'S'{
			res++
		} else if t[i] == 'T'{
			res++
		} else if t[i] == 'Z'{
			res++
		}
	}

	return res // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
}
