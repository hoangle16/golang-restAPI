package user

import "restful_api/entities"

// Repository interface
type Repository interface {
	FetchAll() ([]*entities.User, error)
	FindByID(id uint) (*entities.User, error)
	Store(u *entities.User) (bool, error)
	Update(u *entities.User) (bool, error)
	Delete(id uint) (bool, error)
}
