package entity

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Name     string
	LastName string
	Age      int32
}

func (u *User) IsValid() error {
	err := u.checkPropertiesMustBeDefined()
	return err
}

func (u *User) checkPropertiesMustBeDefined() error {
	if u.Name == "" {
		return errors.New("Name must be defined")
	}

	if u.LastName == "" {
		return errors.New("Last Name must be defined")
	}
	var oldestPersonInTheWorldAge int32 = 120
	if u.Age <= 0 || u.Age > oldestPersonInTheWorldAge {
		return errors.New("Age must be greater than zero and less then 120")
	}
	
	return nil
}
