package loader

import (
	"app/internal"
	"encoding/json"
	"fmt"
	"os"
)

type ProductsJSONFile struct {
	file *os.File
}

func NewProductsJSONFile(file *os.File) *ProductsJSONFile {
	return &ProductsJSONFile{
		file: file,
	}
}

type ProductJSON struct {
	Id          int     `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (c *ProductsJSONFile) Load() (products []internal.Product, err error) {
	// function to read products from products.json file and create a slice of products

	var prod []ProductJSON
	err = json.NewDecoder(c.file).Decode(&prod)
	if err != nil {
		fmt.Println("error decoding file: ", c.file)
		return
	}

	for _, v := range prod {
		products = append(products, internal.Product{
			Id: v.Id,
			ProductAttributes: internal.ProductAttributes{
				Description: v.Description,
				Price:       v.Price,
			},
		})
	}
	return
}
