package models

type Task struct {
	Id       uint   `json:"id"`
	Mensagem string `json:"mensagem"`
	Status   int    `json:"status"`
}
