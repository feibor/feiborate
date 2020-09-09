package feiborate

import (
	"fmt"
	"net/http"
)

// ErrorCode 错误编码
type ErrorCode string

const (
	// Fail 请求失败
	Fail = "0"
	// Success 成功
	Success     = "1"
	// NoRecordErr 无记录的错误，主要用于对接gorm的 no record error
	NoRecordErr = "000000"
)

// ValueStruct 错误码的结果集
type ValueStruct struct {
	Code string
	Msg  string
}

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

// NewErrRecordNotFoundF 新增无记录的错误
func NewErrRecordNotFoundF(errMsg string, args ...interface{}) *RespError {
	return NewCommonRespCodeErrorF(NoRecordErr, errMsg, args ...)
}

// IsRecordNotFoundErr 是否是无记录的err
func IsRecordNotFoundErr(err error) bool {
	if err, ok := err.(*RespError); ok {
		if err.ErrorCode == NoRecordErr {
			return true
		}
	}
	return false
}

// NewCommonRespError 创建一个common error
func NewCommonRespError(errMsg string) *RespError {
	return NewRespError(Fail, errMsg)
}

// NewCommonRespErrorF 创建一个common error
func NewCommonRespErrorF(errMsg string, args ...interface{}) *RespError {
	return NewRespError(Fail, fmt.Sprintf(errMsg, args...))
}

// NewCommonRespCodeErrorF 创建一个common error
func NewCommonRespCodeErrorF(code string, errMsg string, args ...interface{}) *RespError {
	return NewRespError(code, fmt.Sprintf(errMsg, args...))
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
