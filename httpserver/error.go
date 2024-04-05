package httpserver

import (
	"net/http"

	"github.com/pkg/errors"
)

type Code int

const (
	CodeOK Code = iota

	CodeInternalServerError
	CodeBadRequest
	CodeUnauthorized
	CodeNotFound
	CodeForbidden
	CodeTimeout

	CodeCreated
	CodeAccepted
	CodeNoContent
	CodeResetContent
)

const (
	CodeMessageOK                  string = "OK"
	CodeMessageInternalServerError string = "Internal Server Error"
	CodeMessageBadRequest          string = "Bad Request"
	CodeMessageUnauthorized        string = "Unauthorized"
	CodeMessageNotFound            string = "Not Found"
	CodeMessageForbidden           string = "Forbidden"
	CodeMessageTimeout             string = "Timeout"
	CodeMessageCreated             string = "Created"
	CodeMessageAccepted            string = "Accepted"
	CodeMessageNoContent           string = "No Content"
	CodeMessageResetContent        string = "Reset Content"
)

var code2MessageM = map[Code]string{
	CodeOK:                  CodeMessageOK,
	CodeInternalServerError: CodeMessageInternalServerError,
	CodeBadRequest:          CodeMessageBadRequest,
	CodeUnauthorized:        CodeMessageUnauthorized,
	CodeNotFound:            CodeMessageNotFound,
	CodeForbidden:           CodeMessageForbidden,
	CodeTimeout:             CodeMessageTimeout,
	CodeCreated:             CodeMessageCreated,
	CodeAccepted:            CodeMessageAccepted,
	CodeNoContent:           CodeMessageNoContent,
	CodeResetContent:        CodeMessageResetContent,
}

type Error struct {
	Status int
	Code   Code
	Err    error
}

var _ error = &Error{}

func (e *Error) Error() string {
	return e.Error()
}

func (e *Error) WithStatus(status int) *Error {
	e.Status = status
	return e
}

func ErrorWithCode(code Code) *Error {
	return &Error{
		Status: http.StatusOK,
		Code:   code,
		Err:    errors.New(code2MessageM[code]),
	}
}

func ErrorWithBadRequest() *Error {
	return &Error{
		Status: http.StatusBadRequest,
		Code:   CodeBadRequest,
		Err:    errors.New(code2MessageM[CodeBadRequest]),
	}
}
func ErrorWithInternalServer() *Error {
	return &Error{
		Status: http.StatusInternalServerError,
		Code:   CodeInternalServerError,
		Err:    errors.New(code2MessageM[CodeInternalServerError]),
	}
}
