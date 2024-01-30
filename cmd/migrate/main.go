package main

import (
	"app/internal/application"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// env
	// ...
	// - database password
	db_pwd := os.Getenv("DATABASE_PASSWORD")

	// app
	// - config
	cfg := &application.ConfigApplicationMigrate{
		Db: &mysql.Config{
			User:   "root",
			Passwd: db_pwd,
			Net:    "tcp",
			Addr:   "localhost:3306",
			DBName: "fantasy_products",
		},
		CustomerFilepath: "./docs/db/json/customers.json",
		ProductFilepath:  "./docs/db/json/products.json",
		InvoiceFilepath:  "./docs/db/json/invoices.json",
		SaleFilepath:     "./docs/db/json/sales.json",
	}
	app := application.NewApplicationMigrate(cfg)
	// - close the app
	defer app.Close()
	// - set up
	if err := app.SetUp(); err != nil {
		fmt.Println(err)
		return
	}
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
