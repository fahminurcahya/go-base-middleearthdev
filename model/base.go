package model

type BaseResponseModel struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Exception string `json:"exception"`
}
