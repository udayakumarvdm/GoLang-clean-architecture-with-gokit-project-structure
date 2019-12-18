package repository

import (
	entity "GoLang_clean_architecture_with_gokit_project_structure/our_project_service/entity"
	vm "GoLang_clean_architecture_with_gokit_project_structure/our_project_service/view_model"
)

type Repository interface {
	ReqFunction(*vm.Request) (*[]entity.Response, error)
}
