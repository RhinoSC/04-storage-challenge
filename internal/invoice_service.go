package internal

// ServiceInvoice is the interface that wraps the basic methods that an invoice service should implement.
type ServiceInvoice interface {
	// FindAll returns all invoices
	FindAll() (i []Invoice, err error)
	// Save saves an invoice
	Save(i *Invoice) (err error)
	// UpdateTotal updates the total of all invoices base on sales and products
	UpdateAllTotal() (err error)
}
