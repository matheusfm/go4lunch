package product

import "github.com/matheusfm/go4lunch/pkg/product"

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) All() ([]*product.Product, error) {
	products := make([]*product.Product, 0, 2)
	products = append(products,
		&product.Product{
			Name:     "PlayStation 4 Slim 1TB Console",
			Category: "Video Games",
			Price:    247.98,
		},
		&product.Product{
			Name:     "Apple MacBook Pro",
			Category: "Electronics",
			Price:    1399.99,
		})
	return products, nil
}
