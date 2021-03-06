package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int
	msg     string
	details []string
}

// 全局错误码字典作为存储载体
var codes = map[int]string{}

// NewError 创建Error
func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

// Error 错误码字符串方法
func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d，错误信息：%s", e.Code(), e.Msg())
}

// Code 获取错误码
func (e *Error) Code() int {
	return e.code
}

// Msg 获取错误码信息
func (e *Error) Msg() string {
	return e.msg
}

// Msg 格式化错误码
func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

// Details 返回详细信息
func (e *Error) Details() []string {
	return e.details
}

// WithDetails 创建具备详细信息的错误码
func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	newError.details = append(newError.details, details...)

	return &newError
}

// StatusCode 错误码转HTTP状态码
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
