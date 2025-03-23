package services

import (
	"SysAuthenticate/internal/models"
	"SysAuthenticate/internal/repository"
	"SysAuthenticate/internal/utils"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type SignupServiceImpl struct {
	r *repository.SignupRepositoryImpl
}

func NewSignupServiceImpl(r *repository.SignupRepositoryImpl) *SignupServiceImpl {
	return &SignupServiceImpl{r: r}
}

func (s *SignupServiceImpl) RegisterNewUser(ctx context.Context, user *models.User) error {
	if user.Email == "" || user.Password == "" || user.Username == "" {
		return errors.New("Email or Password or Username is empty")
	}

	if len(user.Password) < 8 {
		return errors.New("Password is too short")
	}

	_, err := s.r.FindUserByUsername(ctx, user.Username)
	if err == nil {
		return errors.New("User already exists with same username")
	}

	if utils.IsValidEmail(user.Email) == false {
		return errors.New("Invalid email!")
	}
	_, err = s.r.FindUserByEmail(ctx, user.Email)
	if err == nil {
		return errors.New("User already exists with same email")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPass)

	return s.r.RegisterNewUser(ctx, user)
}
