package transport

import (
	"context"
	"net/http"

	request "github.com/aditiapratama1231/shipping-microservice/pkg/request/http"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/aditiapratama1231/shipping-microservice/pkg/endpoint"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	shipping := r.PathPrefix("/api/shippings").Subrouter()
	shipping.Use(commonMiddleware)

	getShippingsHandler := httptransport.NewServer(
		endpoints.GetShippingsEndpoint,
		request.DecodeGetShippingResponse,
		request.EncodeResponse,
	)

	createNewShippingHandler := httptransport.NewServer(
		endpoints.CreateNewShippingEndpoint,
		request.DecodeCreateShippingRequest,
		request.EncodeResponse,
	)

	shipping.Handle("", getShippingsHandler).Methods("GET")
	shipping.Handle("/create", createNewShippingHandler).Methods("POST")

	return shipping
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
