package routers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/matheusfm/go4lunch/api/handlers"
	"net/http"
)

type SystemRoutes struct {
	productHandler *handlers.Product
}

func NewSystemRoutes(productHandler *handlers.Product) *SystemRoutes {
	return &SystemRoutes{productHandler: productHandler}
}

func (r *SystemRoutes) Routers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
	router.HandleFunc("/products", r.productHandler.AllProducts).Methods(http.MethodGet)
	router.HandleFunc("/products", r.productHandler.CreateProduct).Methods(http.MethodPost)

	router.Use(JsonContentType)
	return router
}

func JsonContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
