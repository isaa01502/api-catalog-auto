package common

import "time"

//Note в рамках /internal/common описаются общие пакеты/модели/утилиты которых можно вызвать откуда угодно

// BaseResponse -
type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type HttpArgs struct {
	Method         string
	Url            string
	Data           []byte
	Headers        map[string]string
	ResponseStruct interface{}
	Proxy          string
	TimeoutSecond  time.Duration
}

type APIError struct {
	ResultCode        int    `json:"resultCode"`
	ResultDescription string `json:"resultDescription"`
}
