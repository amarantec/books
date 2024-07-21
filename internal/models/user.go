package models

type User struct {
	Id       int64
	Name     string `json:"name`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserId   int64  `json:"user_id"`
}
