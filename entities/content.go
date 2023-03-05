package entities

import "time"

type (
	Content struct {
		Id        string    `json:"id"`
		Title     string    `json:"title"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"created_at"`
		UpdateAt  time.Time `json:"update_at"`
	}

	Resp struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	ContentRepo interface {
		GetAll() ([]Content, error)
		Get(id string) (Content, error)
		Save(data Content) error
		Update(data Content) error
		Delete(id string) error
	}

	ContentUseCase interface {
		GetAll() Resp
		Get(id string) Resp
		Save(data Content) Resp
		Update(data Content) Resp
		Delete(id string) Resp
	}
)
