package controller

type Response struct {
	Message string `json:"message"`
}

func NewResponse(msg string) Response {
	return Response{
		Message: msg,
	}
}
