package models

type CreateUser struct {
	Name string
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GetAllUserRequest struct {
	Page  int
	Limit int
	Name  string
}
type GetAllUser struct {
	Users []User
	Count int
}
