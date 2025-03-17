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

func (repo *SignupRepositoryImpl) UpdateUser(ctx context.Context, user *models.User) error {
	return repo.db.WithContext(ctx).Model(&user).Where("id = ?", user.Id).Updates(user).Error
}

func (repo *SignupRepositoryImpl) FindUserByUsername(ctx context.Context, username string) (models.User, error) {
	var user models.User
	if err := repo.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (repo *SignupRepositoryImpl) FindUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	if err := repo.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
