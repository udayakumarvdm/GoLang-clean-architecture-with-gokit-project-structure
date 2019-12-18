package service

import (
	usecase "GoLang_clean_architecture_with_gokit_project_structure/our_project_service/usecase"

	logging "github.com/go-kit/kit/log"
)

type ServiceStruct struct {
	Usecase usecase.ReqUsecase
	log     logging.Logger
}

func NewServiceStructFuction(pcuc usecase.ReqUsecase, log logging.Logger) Service {
	return &ServiceStruct{
		pcuc,
		log,
	}
}
