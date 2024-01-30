package loader

import (
	"app/internal"
	"encoding/json"
	"fmt"
	"os"
)

type CustomersJSONFile struct {
	file *os.File
}

func NewCustomersJSONFile(file *os.File) *CustomersJSONFile {
	return &CustomersJSONFile{
		file: file,
	}
}

type CustomerJSON struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Condition int    `json:"condition"`
}

func (c *CustomersJSONFile) Load() (customers []internal.Customer, err error) {
	// function to read customers from customers.json file and create a slice of customers

	var cus []CustomerJSON
	err = json.NewDecoder(c.file).Decode(&cus)
	if err != nil {
		fmt.Println("error decoding file: ", c.file)
		return
	}

	cus = make([]CustomerJSON, len(cus))
	for _, v := range cus {
		customers = append(customers, internal.Customer{
			Id: v.Id,
			CustomerAttributes: internal.CustomerAttributes{
				FirstName: v.FirstName,
				LastName:  v.LastName,
				Condition: v.Condition,
			},
		})
	}
	return
}
