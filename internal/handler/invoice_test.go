package handler_test

import (
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

// init registers txdb
func init() {
	// db connection
	cfg := mysql.Config{
		User:   "root",
		Passwd: os.Getenv("DATABASE_PASSWORD"),
		Addr:   "127.0.0.1:3306",
		Net:    "tcp",
		DBName: "fantasy_products_test_db",
	}
	// register txdb
	txdb.Register("txdb", "mysql", cfg.FormatDSN())
}

// TestInvoicesDefault_UpdateAllTotal tests the UpdateAllTotal method of the handler.
func TestInvoicesDefault_UpdateAllTotal(t *testing.T) {
	t.Run("case 01: success - updates all invoices total", func(t *testing.T) {
		// arrange
		// ...
		db, err := sql.Open("txdb", "")
		require.NoError(t, err)
		defer db.Close()

		// - database: tear-down
		defer func(db *sql.DB) {
			// delete records
			_, err := db.Exec("DELETE FROM sales")
			if err != nil {
				panic(err)
			}
			_, err = db.Exec("DELETE FROM invoices")
			if err != nil {
				panic(err)
			}
			_, err = db.Exec("DELETE FROM products")
			if err != nil {
				panic(err)
			}
			// reset auto increment
			_, err = db.Exec("ALTER TABLE sales AUTO_INCREMENT = 1")
			if err != nil {
				panic(err)
			}
			_, err = db.Exec("ALTER TABLE invoices AUTO_INCREMENT = 1")
			if err != nil {
				panic(err)
			}
			_, err = db.Exec("ALTER TABLE products AUTO_INCREMENT = 1")
			if err != nil {
				panic(err)
			}
		}(db)

		// - database: set-up
		err = func(db *sql.DB) error {
			// insert invoices
			_, err := db.Exec(
				"INSERT INTO invoices (`id`, `total`) VALUES" +
					"(1, 0)," +
					"(2, 0)," +
					"(3, 0);",
			)
			if err != nil {
				return err
			}
			// insert products
			_, err = db.Exec(
				"INSERT INTO products (`id`, `price`) VALUES" +
					"(1, 10)," +
					"(2, 20)," +
					"(3, 30);",
			)
			if err != nil {
				return err
			}
			// insert sales
			_, err = db.Exec(
				"INSERT INTO sales (`id`, `invoice_id`, `product_id`, `quantity`) VALUES" +
					"(1, 1, 1, 1)," +
					"(2, 1, 2, 1)," +
					"(3, 2, 2, 1)," +
					"(4, 2, 3, 1)," +
					"(5, 3, 3, 1);",
			)
			if err != nil {
				return err
			}
			return nil
		}(db)
		require.NoError(t, err)

		// - repository: mysql
		rp := repository.NewInvoicesMySQL(db)
		// - service: default
		sv := service.NewInvoicesDefault(rp)
		// - handler
		hd := handler.NewInvoicesDefault(sv)
		hdFunc := hd.UpdateAllTotal()

		// act
		request := httptest.NewRequest(http.MethodPut, "/invoices/total", nil)
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"message": "invoices updated"}`
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
	})
}
