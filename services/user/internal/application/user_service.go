package application

import (
	"github.com/ilyasa1211/ecommerce-scale-server/backend/services/user/internal/domain"
	"github.com/ilyasa1211/ecommerce-scale-server/backend/services/user/internal/ports/outbound"
)

type UserService struct {
	repo outbound.UserRepository
}

func NewUserService(repo outbound.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(user *domain.User) error {
	exists, err := s.repo.ExistsByEmail(user.Email)
	if err != nil {
		return err
	}

	if exists {
		return domain.ErrUserAlreadyExists
	}

	return s.repo.Save(user)
}
