package utils

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ErrorRequestInvaild ..
const (
	ErrorRequestInvaild      = 101
	ErrorRequestUnauthorized = 102
)

// ResponseMessage ..
type ResponseMessage struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success,omitempty"`
}

// ErrorMessage ..
type ErrorMessage struct {
	StatusCode int    `json:"-"`
	ErrorCode  int    `json:"error_code,omitempty"`
	Error      string `json:"error_message,omitempty"`
	Success    bool   `json:"success,omitempty"`
	IsDebug    bool   `json:"-"`
	IsAlerts   bool   `json:"-"`
}

// SuccessResponseMessage ..
func SuccessResponseMessage(statusCode int, data interface{}, c echo.Context) error {
	return c.JSON(statusCode, ResponseMessage{
		Data:    data,
		Success: true,
	})
}

// HandleError ..
func HandleError(errorMessage ErrorMessage) *ErrorMessage {
	if errorMessage.IsAlerts {
		//TODO : send notify (line, slack, etc...)
	}
	if errorMessage.IsDebug {
		fmt.Println("[DEBUG] ====> ", errorMessage.Error)
	}

	switch errorMessage.StatusCode {
	case http.StatusInternalServerError:
		errorMessage.Error = "Internal Server Error"
		errorMessage.Success = false
	case http.StatusBadGateway:
		errorMessage.Error = "Bad Gateway Error"
		errorMessage.Success = false
	}

	return &errorMessage
}

// BadRequestError ..
func BadRequestError(c echo.Context, err error) error {
	if err == nil {
		err = fmt.Errorf("Error Request Invaild")
	}
	return c.JSON(http.StatusBadRequest, ErrorMessage{
		StatusCode: http.StatusBadRequest,
		ErrorCode:  ErrorRequestInvaild,
		Error:      err.Error(),
		Success:    false,
	})
}

// UnauthorizedRequestError ..
func UnauthorizedRequestError(c echo.Context, err error) error {
	if err == nil {
		err = fmt.Errorf("Authenication failed.")
	}
	return c.JSON(http.StatusUnauthorized, ErrorMessage{
		StatusCode: http.StatusUnauthorized,
		ErrorCode:  ErrorRequestUnauthorized,
		Error:      err.Error(),
		Success:    false,
	})
}
