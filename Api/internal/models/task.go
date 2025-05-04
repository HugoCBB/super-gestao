package models

type Task struct {
	Id       uint   `json:"id"`
	Cliente  string `json:"cliente"`
	Mensagem string `json:"mensagem"`
}
