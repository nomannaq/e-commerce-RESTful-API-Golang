package types

import "time"

type UserStore interface {
	GetUsersByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) (*User, error)
}
type mockUserStore struct {
}

func GetUsersByEmail(email string) (User, error) {
	return User{}, nil
}

type User struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	ID        string    `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
type RegisterUserPayload struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
