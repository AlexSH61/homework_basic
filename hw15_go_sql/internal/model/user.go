package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	PassWord string `json:"password"`
}

func NewUser(id int, name, email, passWord string) *User {
	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		PassWord: passWord,
	}
}
