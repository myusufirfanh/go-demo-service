package error

import (
	Errors "github.com/pkg/errors"
)

type QoalaError struct {
	Code      int
	ErrorCode string
	Message   string
	Status    string
}

// Error returns error type as a string
func (q *QoalaError) Error() string {
	return q.Message
}

// New returnns new error message in standard pkg errors new
func New(msg string) error {
	return Errors.New(msg)
}

// Wrap returns a new error that adds context to the original error
func Wrap(code int, errorCode string, err error, msg string, status string) error {
	return Errors.Wrap(&QoalaError{
		Code:      code,
		ErrorCode: errorCode,
		Message:   msg,
		Status:    status,
	}, err.Error())
}

func (q QoalaError) GetErrorCode() string {
	return q.ErrorCode
}

func (q QoalaError) GetCode() int {
	return q.Code
}

func (q QoalaError) GetMessage() string {
	return q.Message
}

func (q QoalaError) GetStatus() string {
	return q.Status
}
