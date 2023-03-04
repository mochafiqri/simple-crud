package content

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
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

var Contents = map[string]Content{}
