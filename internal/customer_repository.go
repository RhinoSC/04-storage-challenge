package internal

// RepositoryCustomer is the interface that wraps the basic methods that a customer repository should implement.
type RepositoryCustomer interface {
	// FindAll returns all customers saved in the database.
	FindAll() (c []Customer, err error)
	// Save saves a customer into the database.
	Save(c *Customer) (err error)

	//FindInvoicesByCondition returns the total from invoices by condition
	FindInvoicesByCondition() (c []CustomerInvoicesByCondition, err error)

	//FindTopCustomersByAmountSpent returns the top 5 customers by amount spent
	FindTopCustomersByAmountSpent() (c []CustomerTopByAmountSpent, err error)
}
