package usecase

import (
	vm "GoLang_clean_architecture_with_gokit_project_structure/our_project_service/view_model"
)

type ReqUsecase interface {
	ReqFunction(*vm.Request) (*vm.Response, error)
}
