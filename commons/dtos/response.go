package dtos

type (
	StandardResponseReq struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
		Meta    interface{} `json:"meta"`
		Error   error       `json:"error"`
	}

	StandardResponse struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
		Meta    interface{} `json:"meta"`
		Error   interface{} `json:"error"`
	}
)
