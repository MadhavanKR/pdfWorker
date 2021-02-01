package apisurface

type ApiErrorResponse struct {
	Error string `json:"error"`
	Status int `json:"httpStatus"`
}
