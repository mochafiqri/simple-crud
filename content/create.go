package content

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

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
