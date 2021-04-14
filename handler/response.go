package handler

type Response struct {
	Code    int
	Message string
}

func ResponseMessage(code int, message string) (int, interface{}) {
	response := Response{
		Code:    code,
		Message: message,
	}

	return code, response
}
