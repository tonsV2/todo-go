package user

type Service interface {
	Signup(email string, password string) (*User, error)
	SignIn(email string, password string) (*User, error)
	FindByEmail(email string) (*User, error)
	FindById(id uint) (*User, error)
}

func ProvideService(repository Repository) Service {
	return &service{
		repository,
	}
}

type service struct {
	repository Repository
}

func (s service) Signup(email string, password string) (*User, error) {
	panic("implement me")
}

func (s service) SignIn(email string, password string) (*User, error) {
	panic("implement me")
}

func (s service) FindByEmail(email string) (*User, error) {
	panic("implement me")
}

func (s service) FindById(id uint) (*User, error) {
	panic("implement me")
}
