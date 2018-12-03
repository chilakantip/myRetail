package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

type apiErr struct {
	Error apiErrDetails `json:"error"`

	rawErr         error
	httpStatusCode int
}

type apiErrDetails struct {
	Code    string   `json:"code,omitempty"`
	Message string   `json:"message"`
	Errors  []string `json:"errors,omitempty"`
}

const (
	jsonIndent = "    "
	errsPrefix = "MY-RETAIL"
)

var apiErrors = map[int]apiErrDetails{
	101: {Message: "failed to parse product info"},
	102: {Message: "product info validation failed"},
	103: {Message: "create product failed"},
	104: {Message: "product id is empty"},
	105: {Message: "no record found"},
	106: {Message: "failed to get product details"},
	107: {Message: "invalid product id"},
	108: {Message: "no record affected"},
	109: {Message: "failed to delete the product info"},
}

func newApiErr(code int, errIn error, errors ...string) *apiErr {
	e := apiErr{
		rawErr: errIn,
	}
	errDetails, ok := apiErrors[code]
	if !ok {
		errDetails = apiErrDetails{Code: "UNKNOWN"}
	}

	e.Error = errDetails
	if e.Error.Code == "" {
		e.Error.Code = fmt.Sprintf("%s-%04d", errsPrefix, code)
	}

	if e.httpStatusCode == 0 {
		e.httpStatusCode = http.StatusInternalServerError
	}
	e.Error.Errors = errors

	msgs := []string{}
	if errDetails.Message != "" {
		msgs = append(msgs, errDetails.Message)
	}
	if errIn != nil {
		msgs = append(msgs, errIn.Error())
	}
	e.Error.Message = strings.Join(msgs, "; ")

	return &e
}

func apiErrProductInfoParseFailed() *apiErr {
	return newApiErr(101, nil)
}
func apiErrValidationFailed(err error) *apiErr {
	return newApiErr(102, err)
}
func apiErrAddProductFailed() *apiErr {
	return newApiErr(103, nil)
}
func apiErrProductIDEmpty() *apiErr {
	return newApiErr(104, nil)
}
func apiErrNoRecord() *apiErr {
	return newApiErr(105, nil)
}
func apiErrGetProductFailed() *apiErr {
	return newApiErr(106, nil)
}
func apiErrInvalidProductID() *apiErr {
	return newApiErr(107, nil)
}
func apiErrFailedUpdateProduct() *apiErr {
	return newApiErr(107, nil)
}
func apiErrNoRecordAffected() *apiErr {
	return newApiErr(108, nil)
}
func apiErrDeleteProductFailed() *apiErr {
	return newApiErr(109, nil)
}

// helpers

func (ae *apiErr) render(c echo.Context) error {
	return c.JSONPretty(ae.httpStatusCode, ae, jsonIndent)
}
