package outputports

import DomainModel "auth/domain/model"

type ExecLoginOutputPort interface {
	ExecLogin(email, password string) (*DomainModel.User, error)
}
