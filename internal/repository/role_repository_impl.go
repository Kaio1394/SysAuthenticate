package repository

import (
	"SysAuthenticate/internal/models"
	"context"
	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
	db *gorm.DB
}

func NewRoleRepositoryImpl(db *gorm.DB) *RoleRepositoryImpl {
	return &RoleRepositoryImpl{db: db}
}

func (r *RoleRepositoryImpl) CreateNewRole(ctx context.Context, role *models.Role) error {
	return r.db.WithContext(ctx).Create(&role).Error
}
