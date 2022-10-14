package main

import (
	"fmt"
)

func MapToSlice(mapData map[string]string) [][]string {
	res := [][]string{}
	for key, val := range mapData {
		arr := []string{key, val}
		res = append(res, arr)
	}
	return res // TODO: replace this
}

func main() {
	mapData := map[string]string{
		"hello": "world", "John": "Doe", "age": "14",
	}
	fmt.Println(MapToSlice(mapData))
}
