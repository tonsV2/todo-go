package user

import "gorm.io/gorm"

type Repository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
	FindById(id uint) (*User, error)
}

func ProvideRepository(db *gorm.DB) Repository {
	return &repository{db}
}

type repository struct {
	db *gorm.DB
}

func (r repository) Create(user *User) error {
	panic("implement me")
}

func (r repository) FindByEmail(email string) (*User, error) {
	panic("implement me")
}

func (r repository) FindById(id uint) (*User, error) {
	panic("implement me")
}
