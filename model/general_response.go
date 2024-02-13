package model

type GeneralResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Count   int         `json:"count"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}
