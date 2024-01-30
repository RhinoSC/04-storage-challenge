package internal

type LoaderSale interface {
	// Load loads sales from a source and returns a slice of sales and an error
	Load() (c []Sale, err error)
}
