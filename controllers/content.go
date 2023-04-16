package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/mochafiqri/simple-crud/commons/dtos"
	"github.com/mochafiqri/simple-crud/commons/entities"
	"github.com/mochafiqri/simple-crud/commons/helper"
	"github.com/mochafiqri/simple-crud/commons/interfaces"
	"net/http"
)

type (
	Resp struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
		Source  string      `json:"source"`
	}

	Handler struct {
		uc interfaces.ContentUseCase
	}
)

func NewHandler(uc interfaces.ContentUseCase) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (h *Handler) Routes(e *echo.Echo) {
	e.GET("/contents", h.ReadAll)

	e.GET("/contents/:id", h.ReadById)
	e.POST("/contents", h.Create)
	e.PUT("/contents/:id", h.Update)
	e.DELETE("/contents/:id", h.Delete)
}

func (h *Handler) Create(e echo.Context) error {
	var (
		req = entities.Content{}
	)
	var err = e.Bind(&req)
	if err != nil {
		return helper.Json(e, dtos.StandardResponseReq{
			Code:  http.StatusBadRequest,
			Error: err,
		})
	}

	code, msg, err := h.uc.Create(&req)
	return helper.Json(e, dtos.StandardResponseReq{
		Code:    code,
		Message: msg,
		Data:    req,
		Error:   err,
	})
}

func (h *Handler) ReadAll(e echo.Context) error {
	data, code, err := h.uc.Read()
	return helper.Json(e, dtos.StandardResponseReq{Code: code, Data: data, Error: err})
}

func (h *Handler) ReadById(e echo.Context) error {
	var (
		id = e.Param("id")
	)
	data, code, err := h.uc.Get(id)
	if err != nil {
		return helper.Json(e, dtos.StandardResponseReq{Code: code, Message: err.Error()})
	}

	return helper.Json(e, dtos.StandardResponseReq{Code: code, Data: data})
}

func (h *Handler) Update(e echo.Context) error {
	var (
		id  = e.Param("id")
		req = entities.Content{}
	)

	var err = e.Bind(&req)
	if err != nil {
		return e.JSON(http.StatusBadRequest, Resp{
			Message: http.StatusText(http.StatusBadRequest)})
	}
	req.Id = id

	code, err := h.uc.Update(&req)
	if err != nil {
		return helper.Json(e, dtos.StandardResponseReq{Code: code, Message: err.Error()})

	}
	return helper.Json(e, dtos.StandardResponseReq{Code: code, Data: req})
}

func (h *Handler) Delete(e echo.Context) error {
	var (
		id = e.Param("id")
	)
	code, err := h.uc.Delete(id)
	if err != nil {
		return helper.Json(e, dtos.StandardResponseReq{Code: code, Message: err.Error()})
	}
	return helper.Json(e, dtos.StandardResponseReq{Code: code})
}
