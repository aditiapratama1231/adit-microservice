package cmd

import (
	"github.com/aditiapratama1231/shipping-microservice/database"
	"github.com/aditiapratama1231/shipping-microservice/pkg/endpoint"
	"github.com/aditiapratama1231/shipping-microservice/pkg/service"
)

var (
	db = database.DBInit()

	srvShipping = service.NewShippingService(db)
	Endpoints   = endpoint.Endpoints{
		GetShippingsEndpoint:      endpoint.MakeGetShippingsEndpoint(srvShipping),
		CreateNewShippingEndpoint: endpoint.MakeCreateNewShippingEndpoint(srvShipping),
	}
)
