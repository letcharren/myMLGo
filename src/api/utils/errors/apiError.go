package errors

type ApiError struct{

	Message string `json:"string"`
	Status int `json:"int"`
}