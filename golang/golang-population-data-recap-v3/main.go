package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]interface{} {
	res := []map[string]interface{}{}

	for i := 0; i < len(data); i++ {
		res = append(res, make(map[string]interface{}))
		arr := strings.Split(data[i], ";")
		name := arr[0]
		age, _ := strconv.Atoi(arr[1])
		address := arr[2]
		res[i]["name"] = name
		res[i]["age"] = age
		res[i]["address"] = address
		if arr[3] != "" {
			height, _ := strconv.ParseFloat(arr[3], 64)
			res[i]["height"] = height
		}
		if arr[4] != "" {
			is_married, _ := strconv.ParseBool(arr[4])
			res[i]["isMarried"] = is_married
		}
	}

	return res // TODO: replace this
}

func main() {
	data := []string{"Budi;23;Jakarta;160.42;true",
		"Joko;30;Bandung;160;true",
		"Susi;25;Bogor;165.42;false",
		"Santi;27;Jakarta;160;true",
		"Rudi;23;Jakarta;161.1;false",
		"Rudi 2;23;Jakarta;166.5;false",
		"Rudi 3;23;Jakarta;170;false"}
	fmt.Println(PopulationData(data))
}
