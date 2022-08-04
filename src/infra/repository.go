package infra

import (
	"fmt"

	"github.com/JoaoDeluchi/clean-architecture/src/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type database struct {
	DB *gorm.DB
}

type repo interface {
	createUser()
	getAllUsers()
	getUserById()
	updateUser()
	deleteUser()
}

func (session *database) getAllUsers() []entity.User {
	var users []entity.User
	if query_result := session.DB.Find(&users); query_result.Error != nil {
		fmt.Println(query_result.Error)
	}

	return users
}

func (session *database) createUser(newUser entity.User) entity.User {
	if query_result := session.DB.Create(&newUser); query_result.Error != nil {
		fmt.Println(query_result.Error)
	}

	return newUser
}

func (session *database) getUserById(id uuid.UUID) entity.User {
	var current_user entity.User
	if query_result := session.DB.First(&current_user, id); query_result.Error != nil {
		fmt.Println(query_result.Error)
	}

	return current_user
}

func (session *database) updateUser(id uuid.UUID) entity.User {
	var current_user entity.User
	if query_result := session.DB.First(&current_user, id); query_result.Error != nil {
		fmt.Println(query_result.Error)
	}

	return current_user
}

func (session *database) deleteUser(id uuid.UUID) {
	var user entity.User
	if query_result := session.DB.First(&user, id); query_result.Error != nil {
		fmt.Println(query_result.Error)
	}
	session.DB.Delete(&user)
}
