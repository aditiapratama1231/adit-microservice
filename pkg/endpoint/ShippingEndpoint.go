package endpoint

import (
	"context"

	payload "github.com/aditiapratama1231/shipping-service/pkg/request/payload"

	"github.com/aditiapratama1231/shipping-service/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

func MakeGetShippingsEndpoint(srv service.ShippingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		d := srv.GetShippings()
		return d, nil
	}
}

func MakeCreateNewShippingEndpoint(srv service.ShippingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.CreateNewShippingRequest)
		d := srv.CreateNewShipping(req)
		return d, nil
	}
}
