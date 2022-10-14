package main

import (
	"fmt"
	"os"
	"sort"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

func RecordTransactions(path string, transactions []Transaction) error {
	data := ""
	newData := ""
	rec := map[string]int{}
	recStr := []string{}

	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].Date < transactions[j].Date
	})

	for i := 0; i < len(transactions); i++ {
		if transactions[i].Type == "expense" {
			rec[transactions[i].Date] -= transactions[i].Amount
		} else if transactions[i].Type == "income" {
			rec[transactions[i].Date] += transactions[i].Amount
		}

	}

	for k := range rec {
		recStr = append(recStr, k)
	}

	sort.Strings(recStr)

	for _, k := range recStr {
		if rec[k] < 0 {
			data += k + ";expense;" + fmt.Sprint(rec[k]*-1) + "\n"
		} else {
			data += k + ";income;" + fmt.Sprint(rec[k]) + "\n"
		}

	}

	for i := 0; i < len(data)-1; i++ {
		newData += string(data[i])
	}

	err := os.WriteFile(path, []byte(newData), 0644)

	return err
}

func main() {
	// bisa digunakan untuk pengujian test case
	var transactions = []Transaction{
		{"01/01/2021", "income", 100000},
		{"01/01/2021", "income", 100000},
		{"01/01/2021", "income", 100000},
		{"01/01/2021", "income", 100000},
	}

	err := RecordTransactions("transactions.txt", transactions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}
