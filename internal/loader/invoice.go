package loader

import (
	"app/internal"
	"encoding/json"
	"fmt"
	"os"
)

type InvoicesJSONFile struct {
	file *os.File
}

func NewInvoicesJSONFile(file *os.File) *InvoicesJSONFile {
	return &InvoicesJSONFile{
		file: file,
	}
}

type InvoiceJSON struct {
	Id         int     `json:"id"`
	Datetime   string  `json:"datetime"`
	Total      float64 `json:"total"`
	CustomerId int     `json:"customer_id"`
}

func (c *InvoicesJSONFile) Load() (invoices []internal.Invoice, err error) {
	// function to read invoices from invoices.json file and create a slice of invoices

	var inv []InvoiceJSON
	err = json.NewDecoder(c.file).Decode(&inv)
	if err != nil {
		fmt.Println("error decoding file: ", c.file)
		return
	}

	for _, v := range inv {
		invoices = append(invoices, internal.Invoice{
			Id: v.Id,
			InvoiceAttributes: internal.InvoiceAttributes{
				Datetime:   v.Datetime,
				Total:      v.Total,
				CustomerId: v.CustomerId,
			},
		})
	}
	return
}
