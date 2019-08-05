package endpoint

import "github.com/go-kit/kit/endpoint"

type Endpoints struct {
	GetShippingsEndpoint      endpoint.Endpoint
	CreateNewShippingEndpoint endpoint.Endpoint
}
