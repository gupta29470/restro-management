package models

type CustomError struct {
	Status_Code int    `json:"status_code"`
	Message     string `json:"message"`
	Error       string `json:"error"`
}
