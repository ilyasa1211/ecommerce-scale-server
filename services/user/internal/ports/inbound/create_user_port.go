package inbound

import "github.com/ilyasa1211/ecommerce-scale-server/backend/services/user/internal/domain"

type CreateUserPort interface {
	CreateUser(user *domain.User) error
}
