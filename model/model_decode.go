package model

import (
	viewmodel "GoLang_clean_architecture_with_gokit_project_structure/our_project_service/view_model"
	"context"
	"encoding/json"
	"net/http"
)

//Decode the Push notification Req.
func DecodeReq(_ context.Context, r *http.Request) (interface{}, error) {
	decode := json.NewDecoder(r.Body)
	var Req viewmodel.Request
	err := decode.Decode(&Req)
	if err != nil {
		return nil, ErrBadRouting
	}
	return Req, nil
}
