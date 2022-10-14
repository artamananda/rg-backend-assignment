package invoice

import (
	"errors"
	"strconv"
	"strings"
)

// Marketing invoice
type MarketingInvoice struct {
	Date        string
	StartDate   string
	EndDate     string
	PricePerDay int
	AnotherFee  int
	Approved    bool
}

func (mi MarketingInvoice) RecordInvoice() (InvoiceData, error) {
	var res InvoiceData
	var err error
	months := []string{"", "January", "February", "March", "April", "May",
		"June", "July", "August", "September", "October",
		"November", "December"}

	if mi.Date == "" {
		err = errors.New("invoice date is empty")
		return res, err
	} else if mi.StartDate == "" || mi.EndDate == "" {
		err = errors.New("travel date is empty")
		return res, err
	} else if mi.PricePerDay <= 0 {
		err = errors.New("price per day is not valid")
		return res, err
	}

	sDate := strings.Split(mi.StartDate, "/")
	eDate := strings.Split(mi.EndDate, "/")
	date := strings.Split(mi.Date, "/")

	s, _ := strconv.Atoi(sDate[0])
	e, _ := strconv.Atoi(eDate[0])
	md, _ := strconv.Atoi(date[1])

	total := float64(e-s+1)*float64(mi.PricePerDay) + float64(mi.AnotherFee)

	res.Date = date[0] + "-" + months[md] + "-" + date[2]
	res.TotalInvoice = total
	res.Departemen = "marketing"

	return res, err
}
