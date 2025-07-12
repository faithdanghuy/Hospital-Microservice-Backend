package response

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ResOk { status: "OK_200", data: {} }
type ResOk struct {
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ResErr { status: "ERR_500", message: "invalid request" }
type ResErr struct {
	Status  string   `json:"status,omitempty"`
	Message string   `json:"message,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}

func Error(c echo.Context, code int, message ...string) error {
	if len(message) == 0 {
		return c.JSON(code, ResErr{
			Status:  fmt.Sprintf("ERR_%d", code),
			Message: http.StatusText(code),
		})
	}
	return c.JSON(code, ResErr{
		Status:  fmt.Sprintf("ERR_%d", code),
		Message: message[0],
	})
}

func Errors(c echo.Context, code int, err error, messages ...string) error {
	var resMsg string
	if len(messages) == 0 {
		resMsg = http.StatusText(code)
	} else {
		resMsg = messages[0]
	}
	var errs = strings.Split(err.Error(), "\n")
	return c.JSON(code, ResErr{
		Status:  fmt.Sprintf("ERR_%d", code),
		Message: resMsg,
		Errors:  errs,
	})
}

func OK(c echo.Context, code int, message string, data interface{}) error {
	return c.JSON(code, ResOk{
		Status:  fmt.Sprintf("OK_%d", code),
		Message: message,
		Data:    data,
	})
}

func SimpleOK(c echo.Context, code int, data interface{}) error {
	return c.JSON(code, ResOk{
		Status:  fmt.Sprintf("OK_%d", code),
		Message: "OK",
		Data:    data,
	})
}
