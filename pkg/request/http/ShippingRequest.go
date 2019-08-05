package request

import (
	"context"
	"encoding/json"
	"net/http"

	payload "github.com/aditiapratama1231/shipping-service/pkg/request/payload"
)

func DecodeGetShippingResponse(ctx context.Context, r *http.Request) (interface{}, error) {
	return payload.GetShippingResponse{}, nil
}

func DecodeCreateShippingRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req payload.CreateNewShippingRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	return req, err
}
