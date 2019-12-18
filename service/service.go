package service

import vm "GoLang_clean_architecture_with_gokit_project_structure/our_project_service/view_model"

type Service interface {
	RequestFuctionService(req *vm.Request) (interface{}, error)
}
