package api

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func ErrorHandler(ctx fiber.Ctx, err error) error {

	if apiError, ok := err.(Error); ok {
		return ctx.Status(apiError.Code).JSON(apiError)
	}
	apiError := NewError(http.StatusInternalServerError, err.Error())
	return ctx.Status(apiError.Code).JSON(apiError)
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}

func NewError(code int, msg string) Error {
	return Error{
		Code:    code,
		Message: msg,
	}
}

func ErrInvalidID() Error {
	return Error{
		Code:    http.StatusBadRequest,
		Message: "invalid ID",
	}
}

func ErrBadRequest() Error {
	return Error{
		Code:    http.StatusBadRequest,
		Message: "bad request",
	}
}

func ErrUnauthorized() Error {
	return Error{
		Code:    http.StatusUnauthorized,
		Message: "unauthorized",
	}
}

func ErrInternalError() Error {
	return Error{
		Code:    http.StatusInternalServerError,
		Message: "internal error",
	}
}
