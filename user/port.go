package user

import (
	"ecommere.com/domain"
	userHandler "ecommere.com/rest/handlers/user"
)

type Service interface {
	userHandler.Service // embedding
}

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
	// List() ([]*User, error)
	// Delete(userID int) error
	// Update(user User) (*User, error)
}
