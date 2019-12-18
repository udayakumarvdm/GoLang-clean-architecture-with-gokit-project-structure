package server

import (
	vm "GoLang_clean_architecture_with_gokit_project_structure/our_project_service/view_model"
	service "GoLang_clean_architecture_with_gokit_project_structure/service"
	"context"

	endpointHttp "github.com/go-kit/kit/endpoint"
)

//Validation
//Find Campaing date
func MakeEndpoint(s service.Service) endpointHttp.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(vm.Request)
		r, err := s.RequestFuctionService(&req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}
