package error

import (
	"errors"
	"fmt"
	"time"

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
func (e ApiError) Error() string {
	return e.Msg
}

// Returns an API error with a 404 status code and a message for type t
func ErrorNotFound(t string) ApiError {
	return ApiError{Status: 404, Msg: fmt.Sprintf("%s not found", t)}
}

// Returns an API error with a 400 status code and a message for type t
func ErrorInvalidRequest(t string) ApiError {
	return ApiError{Status: 400, Msg: fmt.Sprintf("Invalid request: %s", t)}
}

// Handles an error thrown by any echo context function
// This is used as echo.HTTPErrorHandler
func HandleEchoError(err error, c echo.Context) {
	e, ok := err.(*ApiError)
	if ok {
		// If the error is an API error, we return the status and message
		c.JSON(e.Status, map[string]any{"error": e.Msg})
		// Log the error to the server log
		c.Logger().Errorf("Error: %s", err)
	} else if err == echo.ErrNotFound {
		c.JSON(404, map[string]any{"error": "Not found"})
	} else if err == echo.ErrUnauthorized {
		c.JSON(401, map[string]any{"error": "Unauthorized"})
	} else if err == echo.ErrForbidden {
		// Handles the rate limiter middleware error
		c.JSON(403, map[string]any{"error": "Forbidden"})
	} else if err == echo.ErrTooManyRequests {
		// Handles the rate limiter middleware error
		c.JSON(429, map[string]any{"error": "Too many requests"})
	} else {
		// This handles any other error as a 500 internal server error
		// Therefore we do not expose any internal error details to the client
		c.JSON(500, map[string]any{"error": "Internal server error"})
	}

}

func NewUnauthorizedError(msg string) *ApiError {
	return &ApiError{
		Status: 401,
		Msg:    fmt.Sprintf("Unauthorized: %s", msg),
	}
}

func NewDaoUnimplementedError() *ApiError {
	return &ApiError{
		Status: 501,
		Msg:    "This DAO function is not yet implemented",
	}
}

func NewApiUnimplementedError() *ApiError {
	return &ApiError{
		Status: 501,
		Msg:    "This api function is not yet implemented",
	}
}

func NewDaoDbError() *ApiError {
	return &ApiError{
		Status: 500,
		Msg:    "Database error",
	}
}

func NewDaoValidationError(referenceObjectName, dataExpected, dataGot string) *ApiError {
	return &ApiError{
		Status: 404,
		Msg:    fmt.Sprintf("Object '%v' invalid. Expected '%v' but got '%v'", referenceObjectName, dataExpected, dataGot),
	}
}

func NewDaoReferenceError(referenceObjectName, referenceKey string) *ApiError {
	return &ApiError{
		Status: 404,
		Msg:    fmt.Sprintf("Reference object '%v' with key '%v' not found", referenceObjectName, referenceKey),
	}
}

func NewDaoReferenceIdError(referenceObjectName string, id uint) *ApiError {
	return NewDaoReferenceError(referenceObjectName, fmt.Sprintf("id: %v", id))
}

func NewDaoReferenceVersionedError(referenceObjectName string, id, version uint) *ApiError {
	return NewDaoReferenceError(referenceObjectName, fmt.Sprintf("id: %v, version: %v", id, version))
}

func NewDaoReferenceTimedError(referenceObjectName string, id uint, startValidity time.Time) *ApiError {
	return NewDaoReferenceError(referenceObjectName, fmt.Sprintf("id: %v, startValidty: %v", id, startValidity))
}

func NewDaoNotExistingError(objectName, referenceKey string) *ApiError {
	return &ApiError{
		Status: 404,
		Msg:    fmt.Sprintf("Object '%v' with key '%v' not found", objectName, referenceKey),
	}
}

func NewClaimMissingError(claim string) *ApiError {
	return &ApiError{
		Status: 403,
		Msg:    fmt.Sprintf("Claim '%v' is required to access this", claim),
	}
}
