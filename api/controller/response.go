package controller

type Response struct {
	Message string      `json:"message"`
	Status  int16       `json:"status"`
	Content interface{} `json:"content"`
}

func NewResponse(msg string, status int16, content ...interface{}) Response {
	return Response{
		Message: msg,
		Status:  status,
		Content: content,
	}
}
