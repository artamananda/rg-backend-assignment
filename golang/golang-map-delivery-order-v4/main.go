package main

import (
	"fmt"
	"strconv"
	"strings"
)

func DeliveryOrder(data []string, day string) map[string]float32 {
	res := map[string]float32{}
	adminCost := 0.1

	if day == "selasa" || day == "kamis" || day == "sabtu" {
		adminCost = 0.05
	}

	for i := 0; i < len(data); i++ {
		strSplit := strings.Split(data[i], ":")
		name := strSplit[0] + "-" + strSplit[1]
		cost, _ := strconv.ParseFloat(strSplit[2], 32)

		if strSplit[3] == "JKT" {
			if day != "minggu" {
				res[name] = float32(cost) + (float32(cost) * float32(adminCost))
			}

		} else if strSplit[3] == "BDG" {
			if day == "rabu" || day == "kamis" || day == "sabtu" {
				res[name] = float32(cost) + (float32(cost) * float32(adminCost))
			}
		} else if strSplit[3] == "BKS" {
			if day == "selasa" || day == "kamis" || day == "jumat" {
				res[name] = float32(cost) + (float32(cost) * float32(adminCost))
			}
		} else if strSplit[3] == "DPK" {
			if day == "senin" || day == "selasa" {
				res[name] = float32(cost) + (float32(cost) * float32(adminCost))
			}
		} else {
			continue
		}
	}

	return res // TODO: replace this
}

func main() {
	data := []string{
		"Budi:Gunawan:10000:JKT",
		"Andi:Sukirman:20000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Budi:Gunawan:50000:DPK",
	}

	day := "sabtu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
