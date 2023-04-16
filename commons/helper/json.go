package helper

import (
	"github.com/labstack/echo/v4"
	"github.com/mochafiqri/simple-crud/commons/dtos"
	"net/http"
)

func Json(c echo.Context, req dtos.StandardResponseReq) error {
	if req.Code > 299 {
	}
	var errResp interface{}
	if req.Error != nil {
		errResp = req.Error.Error()
	}
	if req.Message == "" {
		req.Message = http.StatusText(req.Code)
	}
	return c.JSON(req.Code, dtos.StandardResponse{
		Code:    req.Code,
		Message: req.Message,
		Data:    req.Data,
		Error:   errResp,
	})
}
