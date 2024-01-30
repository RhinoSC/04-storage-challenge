package internal

type LoaderProduct interface {
	// Load loads products from a source and returns a slice of products and an error
	Load() (c []Product, err error)
}
