package service

import (
	vm "GoLang_clean_architecture_with_gokit_project_structure/our_project_service/view_model"
	"time"

	er "GoLang_clean_architecture_with_gokit_project_structure/helper"

	"github.com/go-kit/kit/log/level"
)

func (s *ServiceStruct) RequestFuctionService(req *vm.Request) (interface{}, error) {
	level.Info(s.log).Log("Service layer", "Start")
	//call usecase
	resp, err := s.Usecase.ReqFunction(req)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, er.ErrNotFoundError
	}

	defer func(begin time.Time) {
		level.Debug(s.log).Log(
			//"token", token[0],
			"function", "Service function",
			"result", resp,
			"took", time.Since(begin),
		)
	}(time.Now())

	level.Info(s.log).Log("Service layer", "End")
	return resp, nil
}
