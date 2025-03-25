package services

import (
	"SysAuthenticate/internal/models"
	"SysAuthenticate/internal/repository"
	"SysAuthenticate/internal/utils"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("seu_segredo_super_secreto")

type LoginServiceImpl struct {
	r  *repository.SignupRepositoryImpl
	ru *repository.UserRoleRepositoryImpl
	rr *repository.RoleRepositoryImpl
}

func NewLoginServiceImpl(r *repository.SignupRepositoryImpl, ru *repository.UserRoleRepositoryImpl, rr *repository.RoleRepositoryImpl) *LoginServiceImpl {
	return &LoginServiceImpl{r: r, ru: ru, rr: rr}
}

func (s *LoginServiceImpl) Signup(ctx context.Context, userLogin *models.UserLogin, originType string) (string, error) {
	if userLogin.Username == "" {
		return "", errors.New("username is empty")
	}
	if userLogin.Password == "" {
		return "", errors.New("password is empty")
	}

	user, err := s.r.FindUserByUsername(ctx, userLogin.Username)
	if err != nil {
		return "", err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password)); err != nil {
		return "Password Invalid!", err
	}

	userRole, err := s.ru.GetUserRoleById(ctx, user.Id)
	if err != nil {
		return "", err
	}

	role, err := s.rr.FindRoleById(ctx, userRole.RoleID)
	if err != nil {
		return "", err
	}

	tokenString, err := utils.GenerateToken(&user, role.Name, originType)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
