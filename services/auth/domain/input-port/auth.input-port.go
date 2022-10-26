package inputports

import DomainModel "auth/domain/model"

type AuthInputPort interface {
	Login(email, password string) (*DomainModel.User, error)
	SignUp(email, fullname, password string) (*DomainModel.User, error)
}
