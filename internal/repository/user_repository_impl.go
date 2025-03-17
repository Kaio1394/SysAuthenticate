package repository

import (
	"SysAuthenticate/internal/models"
	"context"
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
