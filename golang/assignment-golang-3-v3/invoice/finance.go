package invoice

import (
	"errors"
	"strconv"
	"strings"
)

// Finance invoice
type FinanceInvoice struct {
	Date     string
	Status   InvoiceStatus // status: "paid", "unpaid"
	Approved bool
	Details  []Detail
}

type InvoiceStatus string

const (
	PAID   InvoiceStatus = "paid"
	UNPAID InvoiceStatus = "unpaid"
)

type Detail struct {
	Description string
	Total       int
}

func (fi FinanceInvoice) RecordInvoice() (InvoiceData, error) {
	var res InvoiceData
	var err error
	months := []string{"", "January", "February", "March", "April", "May",
		"June", "July", "August", "September", "October",
		"November", "December"}
	b := true
	total := 0

	for i := 0; i < len(fi.Details); i++ {
		if fi.Details[i].Total <= 0 {
			b = false
		}
		total += fi.Details[i].Total
	}

	if fi.Date == "" {
		err = errors.New("invoice date is empty")
		return res, err
	} else if len(fi.Details) == 0 {
		err = errors.New("invoice details is empty")
		return res, err
	} else if fi.Status != "paid" && fi.Status != "unpaid" {
		err = errors.New("invoice status is not valid")
		return res, err
	} else if !b {
		err = errors.New("total price is not valid")
		return res, err
	}

	date := strings.Split(fi.Date, "/")
	md, _ := strconv.Atoi(date[1])

	res.Date = date[0] + "-" + months[md] + "-" + date[2]
	res.TotalInvoice = float64(total)
	res.Departemen = "finance"

	return res, err // TODO: replace this
}
