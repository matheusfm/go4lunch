package product

type Product struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float32 `json:"price"`
}
