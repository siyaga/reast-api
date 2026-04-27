package repository

import (
	"reast-api/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser inserts into both tables automatically because of GORM associations
func (r *UserRepository) CreateUser(credential *models.UserCredential) error {
	return r.db.Create(credential).Error
}

// FindByIdentifier searches by Email OR Username
func (r *UserRepository) FindByIdentifier(identifier string) (*models.UserCredential, error) {
	var user models.UserCredential
	// Preload fetches the associated Profile data alongside the credentials
	err := r.db.Preload("Profile").
		Where("email = ? OR username = ?", identifier, identifier).
		First(&user).Error
	return &user, err
}
