package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type RowData struct {
	RankWebsite int
	Domain      string
	TLD         string
	IDN_TLD     string
	Valid       bool
	RefIPs      int
}

func ProcessGetTLD(website RowData, ch chan RowData, chErr chan error) {
	if website.Domain == "" {
		chErr <- errors.New("domain name is empty")
	} else if !website.Valid {
		chErr <- errors.New("domain not valid")
	} else if website.RefIPs == -1 {
		chErr <- errors.New("domain RefIPs not valid")
	} else {
		str := strings.Split(website.Domain, ".")
		if str[len(str)-1] == "com" {
			website.TLD = ".com"
			website.IDN_TLD = ".co.id"
		} else if str[len(str)-1] == "org" {
			website.TLD = ".org"
			website.IDN_TLD = ".org.id"
		} else if str[len(str)-1] == "gov" {
			website.TLD = ".gov"
			website.IDN_TLD = ".go.id"
		} else {
			website.TLD = "." + str[len(str)-1]
			website.IDN_TLD = website.TLD
		}

		ch <- website
		chErr <- nil
	}
	time.Sleep(200 * time.Millisecond)
}

// Gunakan variable ini sebagai goroutine di fungsi FilterAndGetDomain
var FuncProcessGetTLD = ProcessGetTLD

func FilterAndFillData(TLD string, data []RowData) ([]RowData, error) {
	res := []RowData{}
	var err error
	ch := make(chan RowData, len(data))
	errCh := make(chan error)

	for _, website := range data {
		//perbaiki code di bawah ini menggunakan goroutine
		go FuncProcessGetTLD(website, ch, errCh)
	}

	time.Sleep(100 * time.Millisecond)
	for i := 0; i < len(data); i++ {
		errTemp := <-errCh
		if errTemp == nil {
			temp := <-ch
			if temp.TLD == TLD {
				res = append(res, temp)
			}
		} else {
			err = errTemp
		}
	}
	return res, err
}

// gunakan untuk melakukan debugging
func main() {
	TLD := ".com"
	data := []RowData{{
		RankWebsite: 1,
		Domain:      "google.com",
		Valid:       true,
		RefIPs:      -1,
		TLD:         "",
		IDN_TLD:     "",
	},
	}
	fmt.Println(FilterAndFillData(TLD, data))
}
