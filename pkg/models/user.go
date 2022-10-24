package models

type User struct {
	Id       int    `json:"-"`
	Email    string `json:"email" binding: "required,email"`
	UserName string `json:"username" binding: "required"`
	Password string `json:"password" binding: "required"`
}
