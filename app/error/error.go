package error

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"
)

var ErrNotImplemented = errors.New("this function is not implemented")

// API error struct
type ApiError struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

// Function to return the error message as string
// Implements the error interface
func (e *ApiError) Error() string {
	return e.Msg
}

type DaoError int // Enum for DAO errors

// Enum values for DAOError
const (
	DAOUndefined DaoError = iota
	DAONotFound
	DAOAlreadyExists
	DOAInvalid
	DAOUnimplemented
)

// Returns an API error with a 404 status code and a message for type t
func ErrorNotFound(t string) ApiError {
	return ApiError{Status: 404, Msg: fmt.Sprintf("%s not found", t)}
}

// Handles an error thrown by any echo context function
// This is used as echo.HTTPErrorHandler
func HandleEchoError(err error, c echo.Context) {
	e, ok := err.(*ApiError)
	if ok {
		c.JSON(e.Status, map[string]any{"error": e.Msg})
		return
	}
	// This handles any other error as a 500 internal server error
	// Therefore we do not expose any internal error details to the client
	c.JSON(500, map[string]any{"error": "Internal server error"})
}

// Handles an error thrown by the DAO
func HandleDAOError(e DaoError, t string) ApiError {
	switch e {
	case DAONotFound:
		return ErrorNotFound(t)
	case DAOAlreadyExists:
		return ApiError{Status: 400, Msg: fmt.Sprintf("%s already exists", t)}
	case DOAInvalid:
		return ApiError{Status: 400, Msg: fmt.Sprintf("Invalid %s", t)}
	case DAOUnimplemented:
		return ApiError{Status: 500, Msg: "DAO function not implemented"}
	default:
		return ApiError{Status: 500, Msg: "Internal DAO error"}
	}
}