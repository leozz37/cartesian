package handler

// Response is the pattern used on response messages
type Response struct {
	Code    int
	Message string
}

// ResponseMessage builds a response message
func ResponseMessage(code int, message string) (int, interface{}) {
	response := Response{
		Code:    code,
		Message: message,
	}

	return code, response
}
