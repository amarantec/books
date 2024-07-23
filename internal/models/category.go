package models

type Category struct {
	Id	int64
	Name	string 	`json:"name"`
	Url	string 	`json:"url"`
}
