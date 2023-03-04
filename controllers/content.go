package controllers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type (
	Content struct {
		Id        string    `json:"id"`
		Title     string    `json:"title"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"created_at"`
		UpdateAt  time.Time `json:"update_at"`
	}

	Resp struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

var Contents = map[string]Content{}

func Create(e echo.Context) error {
	var (
		req = Content{}
	)
	var err = e.Bind(&req)
	if err != nil {
		return e.JSON(http.StatusBadRequest, Resp{
			Message: http.StatusText(http.StatusBadRequest),
			Data:    nil,
		})
	}

	if req.Title == "" || req.Body == "" {
		return e.JSON(http.StatusBadRequest, Resp{
			Message: "title dan body is required",
			Data:    nil,
		})
	}

	//create id using google uuid
	req.Id = uuid.NewString()
	req.CreatedAt = time.Now()
	Contents[req.Id] = req

	return e.JSON(http.StatusOK, Resp{
		Message: http.StatusText(http.StatusOK),
		Data:    req,
	})
}

func Delete(e echo.Context) error {
	var id = e.Param("id")

	//find content
	if _, ok := Contents[id]; !ok {
		return e.JSON(http.StatusBadRequest, Resp{
			Message: "controllers not found",
			Data:    nil,
		})
	}

	delete(Contents, id)

	return e.JSON(http.StatusOK, Resp{
		Message: http.StatusText(http.StatusOK),
		Data:    nil,
	})
}

func ReadAll(e echo.Context) error {
	var (
		contents = make([]Content, 0)
	)

	if len(Contents) == 0 {
		return e.JSON(http.StatusNotFound, Resp{
			Message: "data not found",
			Data:    nil,
		})
	}

	for _, v := range Contents {
		contents = append(contents, v)
	}

	return e.JSON(http.StatusOK, Resp{
		Message: http.StatusText(http.StatusOK),
		Data:    contents,
	})
}

func ReadById(e echo.Context) error {
	var id = e.Param("id")

	if _, ok := Contents[id]; !ok { // if content not found
		return e.JSON(http.StatusNotFound, Resp{
			Message: "data not found",
			Data:    nil,
		})
	}

	return e.JSON(http.StatusOK, Resp{
		Message: http.StatusText(http.StatusOK),
		Data:    Contents[id],
	})

}

func Update(e echo.Context) error {
	var (
		id      = e.Param("id")
		req     = Content{}
		content = Content{}
	)

	var err = e.Bind(&req)
	if err != nil {
		return e.JSON(http.StatusBadRequest, Resp{
			Message: http.StatusText(http.StatusBadRequest),
			Data:    nil,
		})
	}

	if req.Title == "" || req.Body == "" {
		return e.JSON(http.StatusBadRequest, Resp{
			Message: "title dan body is required",
			Data:    nil,
		})
	}

	//find content
	if _, ok := Contents[id]; !ok {
		return e.JSON(http.StatusBadRequest, Resp{
			Message: "controllers not found",
			Data:    nil,
		})
	}

	content.Id = id
	content.Title = req.Title
	content.Body = req.Body
	content.CreatedAt = Contents[id].CreatedAt
	content.UpdateAt = time.Now()
	Contents[id] = content

	return e.JSON(http.StatusOK, Resp{
		Message: http.StatusText(http.StatusOK),
		Data:    content,
	})
}
