package product

import (
	"github.com/google/wire"
	"github.com/matheusfm/go4lunch/pkg/product"
)

var Providers = wire.NewSet(NewRepository, wire.Bind(new(product.Viewer), new(*Repository)))
