package loader

import (
	"app/internal"
	"encoding/json"
	"fmt"
	"os"
)

type SalesJSONFile struct {
	file *os.File
}

func NewSalesJSONFile(file *os.File) *SalesJSONFile {
	return &SalesJSONFile{
		file: file,
	}
}

type SaleJSON struct {
	Id        int `json:"id"`
	Quantity  int `json:"quantity"`
	ProductId int `json:"product_id"`
	InvoiceId int `json:"invoice_id"`
}

func (c *SalesJSONFile) Load() (sales []internal.Sale, err error) {
	// function to read sales from sales.json file and create a slice of sales

	var sl []SaleJSON
	err = json.NewDecoder(c.file).Decode(&sl)
	if err != nil {
		fmt.Println("error decoding file: ", c.file)
		return
	}

	for _, v := range sl {
		sales = append(sales, internal.Sale{
			Id: v.Id,
			SaleAttributes: internal.SaleAttributes{
				Quantity:  v.Quantity,
				ProductId: v.ProductId,
				InvoiceId: v.InvoiceId,
			},
		})
	}
	return
}
