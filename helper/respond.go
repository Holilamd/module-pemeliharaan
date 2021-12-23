package helper

type Response struct {
	Message interface{} `json:"message"`
	Code    interface{} `json:"status_code"`
	Status  interface{} `json:"status"`
	Data    interface{} `json:"data"`
}

func ResponseSuccess(message string, code int, data interface{}) Response {
	jsonResponse := Response{
		Message: message,
		Code:    code,
		Status:  "success",
		Data:    data,
	}
	return jsonResponse
}

func ResponseError(message string, code int, data interface{}) Response {
	jsonResponse := Response{
		Message: message,
		Code:    code,
		Status:  "Error",
		Data:    data,
	}
	return jsonResponse

}
