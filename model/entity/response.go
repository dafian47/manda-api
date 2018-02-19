package entity

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Paging struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Page int `json:"page"`
	Count int `json:"count"`
	Data    interface{} `json:"data"`
}
