package main

import (
	"errors"
	"fmt"
)

type RowData struct {
	RankWebsite int
	Domain      string
	TLD         string
	IDN_TLD     string
	Valid       bool
	RefIPs      int
}

func GetTLD(domain string) (TLD string, IDN_TLD string) {
	var ListIDN_TLD = map[string]string{
		".com": ".co.id",
		".org": ".org.id",
		".gov": ".go.id",
	}

	for i := len(domain) - 1; i >= 0; i-- {
		if domain[i] == '.' {
			TLD = domain[i:]
			break
		}
	}

	if idnTLD, ok := ListIDN_TLD[TLD]; ok {
		return TLD, idnTLD
	} else {
		return TLD, TLD
	}
}

func ProcessGetTLD(website RowData, ch chan RowData, chErr chan error) {
	if website.Domain == "" {
		chErr <- errors.New("domain name is empty")
		return
	}

	if !website.Valid {
		chErr <- errors.New("domain not valid")
		return
	}

	if website.RefIPs == -1 {
		chErr <- errors.New("domain RefIPs not valid")
		return
	}

	TLD, IDN_TLD := GetTLD(website.Domain)
	website.TLD = TLD
	website.IDN_TLD = IDN_TLD
	ch <- website
}

// Gunakan variable ini sebagai goroutine di fungsi FilterAndGetDomain
var FuncProcessGetTLD = ProcessGetTLD

func FilterAndFillData(TLD string, data []RowData) ([]RowData, error) {
	ch := make(chan RowData, len(data))
	errCh := make(chan error, len(data))

	for _, website := range data {
		go FuncProcessGetTLD(website, ch, errCh)
	}

	var rows []RowData
	var errs []error

	for i := 0; i < len(data); i++ {
		select {
		case row := <-ch:
			rows = append(rows, row)
		case err := <-errCh:
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		var errStr string
		for _, e := range errs {
			errStr += e.Error() + "; "
		}
		return nil, errors.New(errStr[:len(errStr)-2])
	}

	var filteredRows []RowData

	for _, row := range rows {
		if row.TLD == TLD {
			filteredRows = append(filteredRows, row)
		}
	}

	return filteredRows, nil
}

// gunakan untuk melakukan debugging
func main() {
	rows, err := FilterAndFillData(".com", []RowData{
		{1, "google.com", "", "", true, 100},
		{2, "facebook.com", "", "", true, 100},
		{3, "golang.org", "", "", true, 100},
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rows)
}