package product

type Viewer interface {
	All() ([]*Product, error)
}
