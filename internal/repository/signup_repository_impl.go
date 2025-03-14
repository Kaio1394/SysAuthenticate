package repository

import (
	"SysAuthenticate/internal/models"
	"context"
	"gorm.io/gorm"
)

type SignupRepositoryImpl struct {
	db *gorm.DB
}

func NewSignupRepositoryImpl(db *gorm.DB) *SignupRepositoryImpl {
	return &SignupRepositoryImpl{db: db}
}

func (repo *SignupRepositoryImpl) RegisterNewUser(ctx context.Context, user *models.User) error {
	return repo.db.WithContext(ctx).Create(&user).Error
}

func (repo *SignupRepositoryImpl) GetUsers(ctx context.Context) ([]models.User, error) {
	var listUsers []models.User
	if err := repo.db.WithContext(ctx).Find(&listUsers).Error; err != nil {
		return listUsers, err
	}
	return listUsers, nil
}

func (repo *SignupRepositoryImpl) UpdateUser(ctx context.Context, user *models.User) error {
	return repo.db.WithContext(ctx).Model(&user).Where("id = ?", user.Id).Updates(user).Error
}
