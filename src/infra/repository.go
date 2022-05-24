package infra

import "github.com/JoaoDeluchi/clean-architecture/src/entity"

type repo interface {
	createUser()
	getAllUsers()
	getUserById()
	updateUser()
}

func (h *handler) createUser(user *entity.User) {

}
