package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Message interface{} `json:"message"`
	Code    interface{} `json:"status_code"`
	Status  interface{} `json:"status"`
	Data    interface{} `json:"data"`
}

func Success(message string, code int, data interface{}) Response {
	jsonResponse := Response{
		Message: message,
		Code:    code,
		Status:  "success",
		Data:    data,
	}
	return jsonResponse
}

func Error(message string, code int, data interface{}) Response {
	jsonResponse := Response{
		Message: message,
		Code:    code,
		Status:  "Error",
		Data:    data,
	}
	return jsonResponse

}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
