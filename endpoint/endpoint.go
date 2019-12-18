package server

import (
	endpointHttp "github.com/go-kit/kit/endpoint"
)

type EndpointStruct struct {
	RequestFunctionEndpoint endpointHttp.Endpoint
}
