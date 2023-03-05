package controllers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mochafiqri/simple-crud/entities"
	"net/http"
	"time"
)

func Create(e echo.Context) error {
	var (
		req = entities.Content{}
	)
	var err = e.Bind(&req)
	if err != nil {
		return e.JSON(http.StatusBadRequest, entities.Resp{
			Message: http.StatusText(http.StatusBadRequest),
			Data:    nil,
		})
	}

	if req.Title == "" || req.Body == "" {
		return e.JSON(http.StatusBadRequest, entities.Resp{
			Message: "title dan body is required",
			Data:    nil,
		})
	}

	//create id using google uuid
	req.Id = uuid.NewString()
	req.CreatedAt = time.Now()
	entities.Contents[req.Id] = req

	return e.JSON(http.StatusOK, entities.Resp{
		Message: http.StatusText(http.StatusOK),
		Data:    req,
	})
}

func Delete(e echo.Context) error {
	var id = e.Param("id")

	//find content
	if _, ok := entities.Contents[id]; !ok {
		return e.JSON(http.StatusBadRequest, entities.Resp{
			Message: "controllers not found",
			Data:    nil,
		})
	}

	delete(entities.Contents, id)

	return e.JSON(http.StatusOK, entities.Resp{
		Message: http.StatusText(http.StatusOK),
		Data:    nil,
	})
}

func ReadAll(e echo.Context) error {
	var (
		contents = make([]entities.Content, 0)
	)

	if len(entities.Contents) == 0 {
		return e.JSON(http.StatusNotFound, entities.Resp{
			Message: "data not found",
			Data:    nil,
		})
	}

	for _, v := range entities.Contents {
		contents = append(contents, v)
	}

	return e.JSON(http.StatusOK, entities.Resp{
		Message: http.StatusText(http.StatusOK),
		Data:    contents,
	})
}

func ReadById(e echo.Context) error {
	var id = e.Param("id")

	if _, ok := entities.Contents[id]; !ok { // if content not found
		return e.JSON(http.StatusNotFound, entities.Resp{
			Message: "data not found",
			Data:    nil,
		})
	}

	return e.JSON(http.StatusOK, entities.Resp{
		Message: http.StatusText(http.StatusOK),
		Data:    entities.Contents[id],
	})

}

func Update(e echo.Context) error {
	var (
		id      = e.Param("id")
		req     = entities.Content{}
		content = entities.Content{}
	)

	var err = e.Bind(&req)
	if err != nil {
		return e.JSON(http.StatusBadRequest, entities.Resp{
			Message: http.StatusText(http.StatusBadRequest),
			Data:    nil,
		})
	}

	if req.Title == "" || req.Body == "" {
		return e.JSON(http.StatusBadRequest, entities.Resp{
			Message: "title dan body is required",
			Data:    nil,
		})
	}

	//find content
	if _, ok := entities.Contents[id]; !ok {
		return e.JSON(http.StatusBadRequest, entities.Resp{
			Message: "controllers not found",
			Data:    nil,
		})
	}

	content.Id = id
	content.Title = req.Title
	content.Body = req.Body
	content.CreatedAt = entities.Contents[id].CreatedAt
	content.UpdateAt = time.Now()
	entities.Contents[id] = content

	return e.JSON(http.StatusOK, entities.Resp{
		Message: http.StatusText(http.StatusOK),
		Data:    content,
	})
}
