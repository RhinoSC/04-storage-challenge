package internal

// ServiceProduct is the interface that wraps the basic Product methods.
type ServiceProduct interface {
	// FindAll returns all products.
	FindAll() (p []Product, err error)
	// Save saves a product.
	Save(p *Product) (err error)
	// FindTop5BySales returns the top 5 products by quantity on sales.
	FindTop5BySales() (p []ProductTop5BySales, err error)
}
