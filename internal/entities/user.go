package entities

type User struct {
	Id       int    `json:"id"`
	Surname  string `json:"surname"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
