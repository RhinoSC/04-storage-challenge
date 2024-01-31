package internal

// ServiceCustomer is the interface that wraps the basic methods that a customer service should implement.
type ServiceCustomer interface {
	// FindAll returns all customers
	FindAll() (c []Customer, err error)
	// Save saves a customer
	Save(c *Customer) (err error)

	//FindInvoicesByCondition returns the total from invoices by condition
	FindInvoicesByCondition() (c []CustomerInvoicesByCondition, err error)

	//FindTopCustomersByAmountSpent returns the top 5 customers by amount spent
	FindTopCustomersByAmountSpent() (c []CustomerTopByAmountSpent, err error)
}
