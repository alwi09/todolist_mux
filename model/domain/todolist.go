package domain

type Todolist struct {
	Id          int
	Title       string
	Description string
	Status      bool
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
