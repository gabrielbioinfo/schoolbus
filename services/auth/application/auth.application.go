package application

import DomainModel "auth/domain/model"

type AuthApplication struct{}

func NewAuthApplication() *AuthApplication {
	return &AuthApplication{}
}

func (a *AuthApplication) Login(email, password string) (*DomainModel.User, error) {
	return &DomainModel.User{
		ID:       "1",
		Email:    email,
		Password: password,
	}, nil
}

func (a *AuthApplication) SignUp(email, fullname, password string) (*DomainModel.User, error) {
	return &DomainModel.User{
		ID:       "1",
		Email:    email,
		Password: password,
	}, nil
}
