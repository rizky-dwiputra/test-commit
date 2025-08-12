package usecase

import (
	"codelabs-backend-fiber/internal/user/domain"
	"codelabs-backend-fiber/pkg/security"
	"fmt"
)

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) GetAll() ([]domain.User, error) {
	return u.repo.FindAll()
}

func (u *userUsecase) GetByID(id uint) (*domain.User, error) {
	return u.repo.FindByID(id)
}

func (u *userUsecase) Create(user *domain.User) error {
	hashed, err := security.HashPassword(user.Password)
    if err != nil {
        return err
    }
    user.Password = hashed

    if user.Role == "" {
        user.Role = domain.RoleUser
    }
	
	return u.repo.Create(user)
}

func (u *userUsecase) Login(email, password string) (*domain.User, error) {
    user, err := u.repo.FindByEmail(email)
    if err != nil {
        return nil, fmt.Errorf("invalid email or password")
    }

    if !security.CheckPasswordHash(password, user.Password) {
        return nil, fmt.Errorf("invalid email or password")
    }

    return user, nil
}
