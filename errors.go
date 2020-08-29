package feiborate

import (
	"fmt"
	"net/http"
)

// ErrorCode 错误编码
type ErrorCode int

// ValueStruct 错误码的结果集
type ValueStruct struct {
	Code string
	Msg  string
}

const (
	// Fail 请求失败
	Fail = "0"
	// Success 成功
	Success = "1"
)

// RespError 返回错误
type RespError struct {
	ErrorCode string
	ErrorMsg  string
	// HTTPStatusCode  http的错误编码
	HTTPStatusCode int
}

func (resErr *RespError) Error() string {
	return fmt.Sprintf("%s,error code is:%s", resErr.ErrorMsg, resErr.ErrorCode)
}

// NewRespError 创建自定义错误信息
func NewRespError(errCode string, errMsg string) *RespError {
	return &RespError{
		ErrorCode: errCode,
		ErrorMsg:  errMsg,
		// 默认正常
		HTTPStatusCode: http.StatusOK,
	}
}

// NewCommonRespError 创建一个common error
func NewCommonRespError(errMsg string) *RespError {
	return NewRespError(Fail, errMsg)
}

// NewCommonRespErrorF 创建一个common error
func NewCommonRespErrorF(errMsg string, args ...interface{}) *RespError {
	return NewRespError(Fail, fmt.Sprintf(errMsg, args...))
}

// NewRespErrHTTPStatus 与NewRespErr方法不同的是需要携带httpstatus的值
func NewRespErrHTTPStatus(errCode string, errMsg string, httpStatusCode int) *RespError {
	// 如果为0则显示为默认200
	if httpStatusCode == 0 {
		httpStatusCode = http.StatusOK
	}
	return &RespError{
		ErrorCode: errCode,
		ErrorMsg:  errMsg,
		// 默认正常
		HTTPStatusCode: httpStatusCode,
	}
}
