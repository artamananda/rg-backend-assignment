package main

type RowData struct {
	RankWebsite int
	Domain      string
	TLD         string
	IDN_TLD     string
	Valid       bool
	RefIPs      int
}

func ProcessGetTLD(website RowData, ch chan RowData, chErr chan error) {
	// TODO: answer here
}

// Gunakan variable ini sebagai goroutine di fungsi FilterAndGetDomain
var FuncProcessGetTLD = ProcessGetTLD

func FilterAndFillData(TLD string, data []RowData) ([]RowData, error) {
	ch := make(chan RowData, len(data))
	errCh := make(chan error)

	for _, website := range data {
		//perbaiki code di bawah ini menggunakan goroutine
		FuncProcessGetTLD(website, ch, errCh) // TODO: replace this

	}

	return nil, nil // TODO: replace this
}

// gunakan untuk melakukan debugging
func main() {

}
