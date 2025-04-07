package postgres

import (
	"github.com/ilyasa1211/ecommerce-scale-server/backend/services/user/internal/domain"
	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) Save(user *domain.User) error {}
func (r *PostgresUserRepository) FindByID(id uint) (*domain.User, error) {
	user := &domain.User{}

	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (r *PostgresUserRepository) FindByEmail(email string) (*domain.User, error) {
	user := &domain.User{}

	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (r *PostgresUserRepository) FindAll() ([]*domain.User, error)         {}
func (r *PostgresUserRepository) Delete(id uint) error                     {}
func (r *PostgresUserRepository) Count() (int64, error)                    {}
func (r *PostgresUserRepository) ExistsByEmail(email string) (bool, error) {}
func (r *PostgresUserRepository) ExistsByID(id uint) (bool, error)         {}
