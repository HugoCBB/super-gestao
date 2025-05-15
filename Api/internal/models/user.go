package models

type User struct {
	Id     uint   `json:"id"`
	Nome   string `json:"nome"`
	Email  string `json:"email"`
	Senha  string `json:"senha"`
	Status int    `json:"status"`
}
