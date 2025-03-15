package common

type HttpResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessResponse(data interface{}, message string) HttpResponse {
	return HttpResponse{
		Code: 1,
		Msg:  message,
		Data: data,
	}
}

func FailureResponse(data interface{}, message string, code int) HttpResponse {
	return HttpResponse{
		Code: code,
		Msg:  message,
		Data: data,
	}
}
