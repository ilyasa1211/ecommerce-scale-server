package postgres

import (
	"github.com/ilyasa1211/ecommerce-scale-server/backend/services/user/internal/domain"
	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) Save(user *domain.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
func (r *PostgresUserRepository) UpdateByID(user *domain.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}
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
func (r *PostgresUserRepository) FindAll(limit, page uint) ([]*domain.User, error) {
	users := []*domain.User{}

	if err := r.db.Offset(int((page - 1) * limit)).Limit(int(limit)).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
func (r *PostgresUserRepository) Delete(id uint) error {
	if err := r.db.Delete(&domain.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
func (r *PostgresUserRepository) Count() (uint, error) {
	var count int64
	if err := r.db.Model(&domain.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return uint(count), nil
}
func (r *PostgresUserRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	if err := r.db.Model(&domain.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
func (r *PostgresUserRepository) ExistsByID(id uint) (bool, error) {
	var count int64
	if err := r.db.Model(&domain.User{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
