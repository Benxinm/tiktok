package myerrors

import "fmt"

type MyError struct {
	ErrorCode int64
	ErrorMsg  string
}

func (e MyError) Error() string {
	return fmt.Sprintf("[%d] %s", e.ErrorCode, e.ErrorMsg)
}

func NewMyError(code int64, msg string) MyError {
	return MyError{
		ErrorCode: code,
		ErrorMsg:  msg,
	}
}
func (e MyError) AddMessage(error string) MyError {
	e.ErrorMsg = error
	return e
}
