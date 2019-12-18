package usecase

import (
	r "GoLang_clean_architecture_with_gokit_project_structure/our_project_service/repository"
	vm "GoLang_clean_architecture_with_gokit_project_structure/our_project_service/view_model"

	logging "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

//call at main
type UsecaseStruct struct {
	repo r.Repository
	log  logging.Logger
}

//Call the function at main
func NewUsecaseStruct(rep r.Repository, log logging.Logger) ReqUsecase {
	return &UsecaseStruct{
		rep,
		log,
	}
}
func (u *UsecaseStruct) ReqFunction(req *vm.Request) (*vm.Response, error) {
	level.Info(u.log).Log("Usecase Request Fuction", "start")
	//bussiness logic

	//Call repo
	resp, err := u.repo.ReqFunction(req)
	if err != nil {
		return nil, err
	}

	var r vm.Response

	r.Code = 202
	r.CodeType = "Success"
	r.Message = "Retrieved List Successfully"
	r.Status = "Ok"
	r.Data = resp
	level.Info(u.log).Log("Usecase Request Fuction", "end")
	return &r, nil
}
