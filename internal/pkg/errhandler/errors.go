package errhandler

import (
	"errors"
	"net/http"
)

var (
	ErrInternalServer = errors.New("internal server error")
)

type CustomError struct {
	Message    string
	StatusCode int
}

func (e CustomError) Error() string {
	return e.Message
}

func MapError(err error) (statusCode int, errMessage string) {
	switch e := err.(type) {
	case CustomError:
		return e.StatusCode, e.Error()
	default:
		return http.StatusInternalServerError, ErrInternalServer.Error()
	}
}
