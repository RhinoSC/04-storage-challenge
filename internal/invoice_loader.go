package internal

type LoaderInvoice interface {
	// Load loads invoices from a source and returns a slice of invoices and an error
	Load() (c []Invoice, err error)
}
