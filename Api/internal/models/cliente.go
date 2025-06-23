package models

type Cliente struct {
	Id     uint   `json:"id"`
	Numero string `json:"numero"`
	Status int    `json:"status"`
}
