//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/matheusfm/go4lunch/api"
	"github.com/matheusfm/go4lunch/api/routers"
	"github.com/matheusfm/go4lunch/internal/product"
)

func SetupApplication() (*routers.SystemRoutes, error) {
	wire.Build(api.Providers, product.Providers)
	return nil, nil
}
