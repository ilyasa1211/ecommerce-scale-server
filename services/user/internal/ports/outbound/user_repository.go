package outbound

import "github.com/ilyasa1211/ecommerce-scale-server/backend/services/user/internal/domain"

type UserRepository interface {
	Save(user *domain.User) error
	Update(id uint, user *domain.User) error
	FindByID(id uint) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindAll() ([]*domain.User, error)
	Delete(id uint) error
	Count() (int64, error)
	ExistsByEmail(email string) (bool, error)
	ExistsByID(id uint) (bool, error)
}
