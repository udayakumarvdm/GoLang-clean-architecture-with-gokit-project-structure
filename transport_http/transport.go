package transport_http

import (
	endpoint "GoLang_clean_architecture_with_gokit_project_structure/endpoint"
	modelHttpMst "GoLang_clean_architecture_with_gokit_project_structure/model"
	"context"
	"net/http"

	"github.com/go-kit/kit/log"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHttpHandler(ctx context.Context, endpoint endpoint.EndpointStruct, logger log.Logger) http.Handler {

	r := mux.NewRouter()
	options := []httpTransport.ServerOption{
		httpTransport.ServerErrorLogger(logger),
		httpTransport.ServerErrorEncoder(modelHttpMst.EncodeError),
	}

	apiV1 := r.PathPrefix("/api/v1").Subrouter()

	//POST /Find Campaing ID
	apiV1.Methods("POST").Path("/example").Handler(httpTransport.NewServer(
		endpoint.RequestFunctionEndpoint,
		modelHttpMst.DecodeReq,
		modelHttpMst.EncodeResponse,
		options...,
	))

	return r

}
