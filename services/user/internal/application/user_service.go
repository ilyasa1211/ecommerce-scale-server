package application

import (
	"github.com/ilyasa1211/ecommerce-scale-server/backend/services/user/internal/adapters/inbound/http/dto"
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

func (s *UserService) Create(user *dto.CreateUserRequest) (*domain.User, error) {
	exists, err := s.repo.ExistsByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, domain.ErrUserAlreadyExists
	}

	newUser := &domain.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	return newUser, s.repo.Save(newUser)
}

func (s *UserService) GetByID(id uint) (*domain.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (s *UserService) GetByEmail(email string) (*domain.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (s *UserService) UpdateByID(id uint, user *dto.UpdateUserRequest) (*domain.User, error) {
	existingUser, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if user.Name != "" {
		existingUser.Name = user.Name
	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}

	return existingUser, s.repo.Save(existingUser)
}
func (s *UserService) DeleteByID(id uint) error {
	return s.repo.Delete(id)
}
func (s *UserService) GetAll(page uint) ([]*domain.User, error) {
	var limit uint = 15
	users, err := s.repo.FindAll(limit, page)
	if err != nil {
		return nil, err
	}

	return users, nil
}
