package controller

import (
	"context"
	"encoding/json"
	"github.com/ToluwaniO/utube/helper"
	"github.com/ToluwaniO/utube/model"
	"github.com/ToluwaniO/utube/repository"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request)  {
	token := helper.GetUserToken(context.Background(), r.Header.Get("Authorization"))
	var user model.User
	json.Unmarshal([]byte(helper.GetBody(r)), &user)
	resp := model.Response{
		Code:400,
		Status:"failure",
		Message:"An error occurred",
	}
	if token != nil || token.UID == user.UID {
		_, err := repository.AddUser(user.UID, &user)
		if err == nil {
			resp = model.Response{
				Code:   200,
				Status: "success",
				Data:   user,
			}
		}
	}
	data, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(500)
	}
	w.Write(data)
}
