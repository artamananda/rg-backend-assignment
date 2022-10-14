package main

import "fmt"

func GraduateStudent(score int, absent int) string {
	var res = "lulus"

	if score >= 70 && absent < 5 {
		res = "lulus"
	} else {
		res = "tidak lulus"
	}
	return res // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GraduateStudent(100, 4))
}
