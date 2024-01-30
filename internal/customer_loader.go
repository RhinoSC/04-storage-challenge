package internal

type LoaderCustomer interface {
	// Load loads customers from a source and returns a slice of customers and an error
	Load() (c []Customer, err error)
}
