package controller

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func NewResponse(msg, status string) Response {
	return Response{
		Message: msg,
		Status:  status,
	}
}
