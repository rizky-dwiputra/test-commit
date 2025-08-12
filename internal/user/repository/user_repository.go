package repository

import (
	"codelabs-backend-fiber/internal/user/domain"
	customError "codelabs-backend-fiber/pkg/error"
	"strings"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepo{db}
}

func (r *userRepo) FindAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepo) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepo) Create(user *domain.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "idx_users_email") {
			return customError.ErrEmailAlreadyExists
		}
		return err
	}
	return nil
}

func (r *userRepo) FindByEmail(email string) (*domain.User, error) {
    var user domain.User
    if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
