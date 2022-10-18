package helper

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(message string, data interface{}) Response {
	res := Response{
		Message: message,
		Data:    data,
	}

	return res
}

func BuildErrorResponse(message string, data interface{}) Response {

	res := Response{
		Message: message,
		Data:    data,
	}

	return res
}
