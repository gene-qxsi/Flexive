package repository

import (
	"context"

	"github.com/gene-qxsi/Flexive/internal/repository/models"
	"gorm.io/gorm"
)

type ProfileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r *ProfileRepository) CreateProfile(ctx context.Context, profile *models.Profile) (*models.Profile, error) {

	// var profile models.Profile
	err := r.db.WithContext(ctx).Debug().Create(&profile).Error
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (r *ProfileRepository) GetProfileByUserID(ctx context.Context, userID int) (*models.Profile, error) {

	var profile models.Profile
	err := r.db.WithContext(ctx).Debug().Where("user_id = ?", userID).First(&profile).Error
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (r *ProfileRepository) GetProfiles(ctx context.Context) ([]models.Profile, error) {

	var profile []models.Profile
	err := r.db.WithContext(ctx).Debug().Find(&profile).Error
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (r *ProfileRepository) UpdateProfileByUserID(ctx context.Context, userID int, values map[string]interface{}) (*models.Profile, error) {

	var profile models.Profile
	err := r.db.WithContext(ctx).Debug().Where("user_id = ?", userID).First(&profile).Error
	if err != nil {
		return nil, err
	}

	err = r.db.WithContext(ctx).Debug().Model(&profile).Updates(values).Error
	if err != nil {
		return nil, err
	}

	return &profile, nil
}
