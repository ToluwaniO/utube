package model

type Response struct {
	Code int64 `json:"code"`
	Status string `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
	//Headers interface{} `json:"headers"`
}
