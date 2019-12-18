package repository

import (
	entity "GoLang_clean_architecture_with_gokit_project_structure/our_project_service/entity"
	vm "GoLang_clean_architecture_with_gokit_project_structure/our_project_service/view_model"

	logging "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jinzhu/gorm"
)

//call at main
type RepoStruct struct {
	DB  *gorm.DB //Call the database
	log logging.Logger
}

//Call the function at main
func NewRepoStruct(conn *gorm.DB, log logging.Logger) Repository {
	return &RepoStruct{
		conn,
		log,
	}
}

func (r *RepoStruct) ReqFunction(req *vm.Request) (*[]entity.Response, error) {
	level.Info(r.log).Log("Repo Request Fuction", "start")
	//Query part
	var resp []entity.Response
	if err := r.DB.Where("city = ?", req.City).Find(&resp).Error; err != nil {
		return nil, err
	}
	level.Info(r.log).Log("Repo Request Fuction", "end")
	return &resp, nil
}
