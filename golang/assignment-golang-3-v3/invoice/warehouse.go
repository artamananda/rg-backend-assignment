package invoice

import (
	"errors"
)

// Warehouse invoice

type WarehouseInvoice struct {
	Date        string
	InvoiceType InvoiceTypeName
	Approved    bool
	Products    []Product
}

type InvoiceTypeName string

const (
	PURCHASE InvoiceTypeName = "purchase"
	SALES    InvoiceTypeName = "sales"
)

type Product struct {
	Name     string
	Unit     int
	Price    float64
	Discount float64
}

func (wi WarehouseInvoice) RecordInvoice() (InvoiceData, error) {
	var res InvoiceData
	var err error
	var total float64
	b := true
	c := true

	for i := 0; i < len(wi.Products); i++ {
		if wi.Products[i].Unit <= 0 {
			b = false
		}
		if wi.Products[i].Price <= 0 {
			c = false
		}
		total += (float64(wi.Products[i].Unit) * wi.Products[i].Price) - (float64(wi.Products[i].Unit) * wi.Products[i].Price * wi.Products[i].Discount)

	}

	res.Date = wi.Date
	res.TotalInvoice = total
	res.Departemen = "warehouse"

	if wi.Date == "" {
		err = errors.New("invoice date is empty")
	} else if wi.InvoiceType != "purchase" && wi.InvoiceType != "sales" {
		err = errors.New("invoice type is not valid")
	} else if len(wi.Products) == 0 {
		err = errors.New("invoice products is empty")
	} else if !b {
		err = errors.New("unit product is not valid")
	} else if !c {
		err = errors.New("price product is not valid")
	}

	return res, err // TODO: replace this
}
