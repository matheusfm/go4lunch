package api

import (
	"github.com/google/wire"
	"github.com/matheusfm/go4lunch/api/handlers"
	"github.com/matheusfm/go4lunch/api/routers"
)

var Providers = wire.NewSet(routers.NewSystemRoutes, handlers.NewProduct)
