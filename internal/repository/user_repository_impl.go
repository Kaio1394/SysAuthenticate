package repository

import (
	"SysAuthenticate/internal/models"
	"context"
	"errors"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (repo *UserRepositoryImpl) GetUsers(ctx context.Context) ([]models.User, error) {
	var listUsers []models.User
	if err := repo.db.WithContext(ctx).Find(&listUsers).Error; err != nil {
		return listUsers, err
	}
	return listUsers, nil
}

func (repo *UserRepositoryImpl) GetUserById(ctx context.Context, userId uint) (*models.User, error) {
	var user models.User
	if err := repo.db.WithContext(ctx).Where("UserID = ?", userId).First(&user).Error; err != nil {
		return nil, errors.New("UserID not found.")
	}

	return &user, nil
}
