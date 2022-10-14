package main

import (
	"a21hc3NpZ25tZW50/invoice"
	"fmt"
	"log"
)

func RecapDataInvoice(data []invoice.Invoice) ([]invoice.InvoiceData, error) {
	var res []invoice.InvoiceData
	var err error

	if len(data) == 0 {
		return res, err
	}

	for i := 0; i < len(data); i++ {
		a := false
		tmp, err2 := data[i].RecordInvoice()
		err = err2
		invData := invoice.InvoiceData{}
		if i > 0 {
			for j := 0; j < len(res); j++ {
				if tmp.Date == res[j].Date && tmp.Departemen == res[j].Departemen {
					fmt.Println("1")
					res[j].TotalInvoice += tmp.TotalInvoice
					a = true
				}
			}
			if !a {
				fmt.Println("2")
				res = append(res, invData)
				res[len(res)-1].Date = tmp.Date
				res[len(res)-1].Departemen = tmp.Departemen
				res[len(res)-1].TotalInvoice = tmp.TotalInvoice
			}
		} else {
			fmt.Println("3")
			res = append(res, invData)
			res[len(res)-1].Date = tmp.Date
			res[len(res)-1].Departemen = tmp.Departemen
			res[len(res)-1].TotalInvoice = tmp.TotalInvoice
		}

	}

	return res, err // TODO: replace this
}

func main() {
	listInvoice := []invoice.Invoice{
		invoice.FinanceInvoice{
			Date:     "01/02/2020",
			Details:  []invoice.Detail{{Description: "pembelian nota", Total: 4000}, {Description: "Pembelian alat tulis", Total: 4000}},
			Status:   invoice.PAID,
			Approved: true,
		},
		invoice.FinanceInvoice{
			Date:     "01/02/2020",
			Details:  []invoice.Detail{{Description: "pembelian nota", Total: 4000}, {Description: "Pembelian alat tulis", Total: 4000}},
			Status:   invoice.PAID,
			Approved: true,
		},
		invoice.WarehouseInvoice{
			Date: "01-February-2020",
			Products: []invoice.Product{
				{Name: "product A", Unit: 10, Price: 10000, Discount: 0.1},
				{Name: "product C", Unit: 5, Price: 15000, Discount: 0.2},
			},
			InvoiceType: invoice.PURCHASE,
			Approved:    true,
		},
		invoice.MarketingInvoice{
			Date:        "01/02/2020",
			StartDate:   "20/01/2020",
			EndDate:     "25/01/2020",
			Approved:    true,
			PricePerDay: 10000,
			AnotherFee:  5000,
		},
	}

	result, err := RecapDataInvoice(listInvoice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
