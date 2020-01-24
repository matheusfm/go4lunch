package handlers

import (
	"encoding/json"
	"github.com/matheusfm/go4lunch/pkg/product"
	"net/http"
)

type Product struct {
	productViewer product.Viewer
}

func NewProduct(productViewer product.Viewer) *Product {
	return &Product{productViewer: productViewer}
}

func (r *Product) AllProducts(w http.ResponseWriter, req *http.Request) {
	products, _ := r.productViewer.All()
	bytes, _ := json.Marshal(products)
	w.Write(bytes)
}

func (r *Product) CreateProduct(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
