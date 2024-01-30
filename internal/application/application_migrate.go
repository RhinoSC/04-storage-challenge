package application

import (
	"app/internal"
	"app/internal/loader"
	"app/internal/migrator"
	"app/internal/repository"
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

// ConfigApplicationDefault is the configuration for NewApplicationDefault.
type ConfigApplicationMigrate struct {
	// Db is the database configuration.
	Db *mysql.Config
	// Customer Filepath
	CustomerFilepath string
	// Product Filepath
	ProductFilepath string
	// Invoice Filepath
	InvoiceFilepath string
	// Sale Filepath
	SaleFilepath string
}

type ApplicationMigrate struct {
	config       *ConfigApplicationMigrate
	database     *sql.DB
	CustomerFile *os.File
	ProductFile  *os.File
	InvoiceFile  *os.File
	SaleFile     *os.File
	migrators    []internal.Migrator
}

func NewApplicationMigrate(config *ConfigApplicationMigrate) *ApplicationMigrate {
	return &ApplicationMigrate{
		config: config,
	}
}

func (a *ApplicationMigrate) Close() {

	// Close the files
	if a.CustomerFile != nil {
		a.CustomerFile.Close()
	}
	if a.ProductFile != nil {
		a.ProductFile.Close()
	}
	if a.InvoiceFile != nil {
		a.InvoiceFile.Close()
	}
	if a.SaleFile != nil {
		a.SaleFile.Close()
	}

	// close the database connection
	a.database.Close()
}

func (a *ApplicationMigrate) SetUp() (err error) {
	a.database, err = sql.Open("mysql", a.config.Db.FormatDSN())
	if err != nil {
		return
	}
	err = a.database.Ping()
	if err != nil {
		return
	}

	// Open the files
	a.CustomerFile, err = os.Open(a.config.CustomerFilepath)
	if err != nil {
		return
	}
	a.ProductFile, err = os.Open(a.config.ProductFilepath)
	if err != nil {
		return
	}
	a.InvoiceFile, err = os.Open(a.config.InvoiceFilepath)
	if err != nil {
		return
	}
	a.SaleFile, err = os.Open(a.config.SaleFilepath)
	if err != nil {
		return
	}

	// create the migrators
	ldCustomer := loader.NewCustomersJSONFile(a.CustomerFile)
	rpCustomer := repository.NewCustomersMySQL(a.database)
	mgCustomer := migrator.NewMigratorCustomer(ldCustomer, rpCustomer)

	ldInvoice := loader.NewInvoicesJSONFile(a.InvoiceFile)
	rpInvoice := repository.NewInvoicesMySQL(a.database)
	mgInvoice := migrator.NewMigratorInvoice(ldInvoice, rpInvoice)

	ldProduct := loader.NewProductsJSONFile(a.ProductFile)
	rpProduct := repository.NewProductsMySQL(a.database)
	mgProduct := migrator.NewMigratorProduct(ldProduct, rpProduct)

	ldSale := loader.NewSalesJSONFile(a.SaleFile)
	rpSale := repository.NewSalesMySQL(a.database)
	mgSale := migrator.NewMigratorSale(ldSale, rpSale)

	a.migrators = []internal.Migrator{
		mgCustomer,
		mgInvoice,
		mgProduct,
		mgSale,
	}
	return
}

func (a *ApplicationMigrate) Run() (err error) {
	for _, v := range a.migrators {
		err = v.Migrate()
		if err != nil {
			return
		}
	}
	return
}
