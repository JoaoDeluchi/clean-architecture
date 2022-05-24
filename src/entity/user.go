package entity

import (
	"errors"

	"github.com/google/uuid"
)

//User data
type User struct {
	ID       uuid.UUID
	Name     string
	LastName string
	Age      int
}

//NewUser create a new user
func NewUser(email, password, name, lastName string, age int) (*User, error) {
	u := &User{
		ID:       uuid.New(),
		Name:     name,
		LastName: lastName,
		Age:      age,
	}

	err := u.Validate()
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) Validate() error {
	if u.Name == "" || u.Age == 0 || u.LastName == "" {
		return errors.New("User invalid")
	}

	return nil
}
