package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
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
		Source  string      `json:"source"`
	}

	Handler struct {
		Db  *gorm.DB
		Rds *redis.Client
	}
)

func (h *Handler) Create(e echo.Context) error {
	var (
		req = Content{}
	)
	var err = e.Bind(&req)
	if err != nil {
		return e.JSON(http.StatusBadRequest, Resp{
			Message: http.StatusText(http.StatusBadRequest)})
	}

	if req.Title == "" || req.Body == "" {
		return e.JSON(http.StatusBadRequest, Resp{
			Message: "title dan body is required"})
	}

	//create id using google uuid
	req.Id = uuid.NewString()

	err = h.Db.Create(&req).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, Resp{
			Message: err.Error()})
	}

	err = h.Rds.Del(context.Background(), "all").Err() //set to redis
	if err != nil {
		log.Println("[ReadAll][Set] err", err.Error())
	}

	return e.JSON(http.StatusOK, Resp{
		Message: http.StatusText(http.StatusOK),
		Data:    req,
	})
}

func (h *Handler) ReadAll(e echo.Context) error {
	var (
		contents      = []Content{}
		codeError int = http.StatusInternalServerError
	)
	var ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	result, err := h.Rds.Get(ctx, "all").Result() //check redis
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return e.JSON(http.StatusInternalServerError, Resp{
				Message: err.Error()})
		}
	}
	if result != "" { // if exist return
		err = json.Unmarshal([]byte(result), &contents)
		if err != nil {
			log.Println("Error [ReadAll][Set] err", err.Error())
		} else {
			return e.JSON(http.StatusOK, Resp{
				Message: http.StatusText(http.StatusOK),
				Data:    contents,
				Source:  "redis",
			})
		}
	}
	err = h.Db.Order("created_at desc").Find(&contents).Error //get DB
	if err != nil {
		return e.JSON(codeError, Resp{Message: err.Error()})
	}
	if len(contents) == 0 { //if not found
		return e.JSON(http.StatusNotFound, Resp{
			Message: http.StatusText(http.StatusNotFound)})
	}
	data, err := json.Marshal(contents)
	if err != nil {
		log.Println("Error [ReadAll][Set] err", err.Error())
	} else {
		err = h.Rds.Set(ctx, "all", data, 3*time.Hour).Err() //set to redis
		if err != nil {
			log.Println("Error [ReadAll][Set] err", err.Error())
		}
	}

	return e.JSON(http.StatusOK, Resp{
		Message: http.StatusText(http.StatusOK),
		Data:    contents,
		Source:  "main database",
	})
}

func (h *Handler) ReadById(e echo.Context) error {
	var (
		id            = e.Param("id")
		content       = Content{}
		codeError int = http.StatusInternalServerError
	)
	var ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	result, err := h.Rds.Get(ctx, id).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			log.Println("Error [ReadById][Get] err ", err.Error())
		}
	}

	if result != "" {
		err = json.Unmarshal([]byte(result), &content)
		if err != nil {
			log.Println("Error [ReadById][Unmarshal] err ", err.Error())
		} else {
			return e.JSON(http.StatusOK, Resp{
				Message: http.StatusText(http.StatusOK),
				Data:    content,
				Source:  "redis",
			})
		}
	}

	err = h.Db.Where("id = ?", id).Find(&content).Error
	if err != nil {
		return e.JSON(codeError, Resp{
			Message: err.Error()})
	}

	if content.Id == "" {
		return e.JSON(http.StatusNotFound, Resp{
			Message: http.StatusText(http.StatusNotFound),
			Data:    nil,
		})
	}

	data, err := json.Marshal(content)
	if err != nil {
		log.Println("Error [ReadById][Marshal] err ", err.Error())
	} else {
		err = h.Rds.Set(ctx, id, data, 3*time.Hour).Err()
		if err != nil {
			log.Println("Error [ReadById][Set] err ", err.Error())
		}
	}

	return e.JSON(http.StatusOK, Resp{
		Message: http.StatusText(http.StatusOK),
		Data:    content,
		Source:  "main database",
	})
}

func (h *Handler) Update(e echo.Context) error {
	var (
		id        = e.Param("id")
		req       = Content{}
		content   = Content{}
		codeError = http.StatusInternalServerError
	)

	var err = e.Bind(&req)
	if err != nil {
		return e.JSON(http.StatusBadRequest, Resp{
			Message: http.StatusText(http.StatusBadRequest)})
	}

	if req.Title == "" || req.Body == "" {
		return e.JSON(http.StatusBadRequest, Resp{
			Message: "title dan body is required"})
	}

	//find content
	err = h.Db.Where("id = ?", id).Find(&content).Error
	if err != nil {
		return e.JSON(codeError, Resp{
			Message: err.Error()})
	}
	if content.Id == "" {
		return e.JSON(http.StatusNotFound, Resp{
			Message: http.StatusText(http.StatusNotFound)})
	}
	content.Id = id
	content.Title = req.Title
	content.Body = req.Body
	content.UpdateAt = time.Now()

	err = h.Db.Updates(&content).Error
	if err != nil {
		return e.JSON(codeError, Resp{Message: err.Error()})
	}

	//delete redis
	err = h.Rds.Del(context.Background(), id).Err()
	if err != nil {
		log.Println("[Redis][Del]", err.Error())
	}

	err = h.Rds.Del(context.Background(), "all").Err()
	if err != nil {
		log.Println("[Redis][Del]", err.Error())
	}

	return e.JSON(http.StatusOK, Resp{
		Message: http.StatusText(http.StatusOK),
		Data:    content,
	})
}

func (h *Handler) Delete(e echo.Context) error {
	var (
		id        = e.Param("id")
		content   = Content{}
		err       error
		codeError int = http.StatusInternalServerError
	)

	//find content
	err = h.Db.Where("id = ?", id).Find(&content).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			codeError = http.StatusBadRequest
		}
		return e.JSON(codeError, Resp{
			Message: err.Error()})
	}

	//delete
	err = h.Db.Where("id = ?", id).Delete(&content).Error
	if err != nil {
		return e.JSON(codeError, Resp{
			Message: err.Error()})
	}

	//delete redis
	err = h.Rds.Del(context.Background(), id).Err()
	if err != nil {
		log.Println("[Redis][Del]", err.Error())
	}

	err = h.Rds.Del(context.Background(), "all").Err()
	if err != nil {
		log.Println("[Redis][Del]", err.Error())
	}

	return e.JSON(http.StatusOK, Resp{
		Message: http.StatusText(http.StatusOK)})
}
