package models

type Task struct {
	Id       uint   `json:"id"`
	Cliente  string `json:"cliente"`
	Mensagem string `json:"mensagem"`
	Numero   string `json:"numero"`
	Status   int    `json:"status"`
}
