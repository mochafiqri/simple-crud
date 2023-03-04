package content

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

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
			Message: "content not found",
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
