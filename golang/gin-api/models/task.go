package models

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Descritpion string `json:"description"`
	Status      string `json:"status"`
}
