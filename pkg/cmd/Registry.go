package cmd

import (
	"github.com/aditiapratama1231/shipping-service/database"
	"github.com/aditiapratama1231/shipping-service/pkg/endpoint"
	"github.com/aditiapratama1231/shipping-service/pkg/service"
)

var (
	db = database.DBInit()

	srvShipping = service.NewShippingService(db)
	Endpoints   = endpoint.Endpoints{
		GetShippingsEndpoint:      endpoint.MakeGetShippingsEndpoint(srvShipping),
		CreateNewShippingEndpoint: endpoint.MakeCreateNewShippingEndpoint(srvShipping),
	}
)
