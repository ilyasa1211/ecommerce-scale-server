package outbound

import "github.com/ilyasa1211/ecommerce-scale-server/backend/services/user/internal/domain"

type UserRepository interface {
	Save(user *domain.User) error
	UpdateByID(user *domain.User) error
	FindByID(id uint) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindAll(limit, page uint) ([]*domain.User, error)
	Delete(id uint) error
	Count() (uint, error)
	ExistsByEmail(email string) (bool, error)
	ExistsByID(id uint) (bool, error)
}
