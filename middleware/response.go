package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type JsonResponse struct {
  Code    int         `json:"code"`    // ((0: Success, 1: Error, >1: Error Code))
	Message string      `json:"message"` 
	Data    interface{} `json:"data"`    
}

func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
    r.Response.WriteJsonExit(JsonResponse{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

func JsonExit(r *ghttp.Request, err int, msg string, data ...interface{}) {
	Json(r, err, msg, data...)
	r.Exit()
}
