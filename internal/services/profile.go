package services

import (
	"context"

	"github.com/gene-qxsi/Flexive/internal/domain"
	"github.com/gene-qxsi/Flexive/internal/repository"
)

type ProfileService struct {
	repo *repository.ProfileRepository
}

func NewProfileService(repo *repository.ProfileRepository) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) CreateProfile(ctx context.Context, profile domain.Profile) (*domain.Profile, error) {

	profileORM := domain.ToORMProfile(&profile)

	profileORM, err := s.repo.CreateProfile(ctx, profileORM)
	if err != nil {
		return nil, err
	}

	return domain.ToDomainProfile(profileORM), nil
}

func (s *ProfileService) GetProfileByUserID(ctx context.Context, userID int) (*domain.Profile, error) {

	profileORM, err := s.repo.GetProfileByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return domain.ToDomainProfile(profileORM), nil
}

func (s *ProfileService) GetProfiles(ctx context.Context) ([]domain.Profile, error) {

	profilesORM, err := s.repo.GetProfiles(ctx)
	if err != nil {
		return nil, err
	}

	var profiles []domain.Profile
	for _, profile := range profilesORM {
		profiles = append(profiles, *domain.ToDomainProfile(&profile))
	}

	return profiles, nil
}

func (s *ProfileService) UpdateProfileByUserID(ctx context.Context, userID int, values map[string]interface{}) (*domain.Profile, error) {

	profileORM, err := s.repo.UpdateProfileByUserID(ctx, userID, values)
	if err != nil {
		return nil, err
	}

	return domain.ToDomainProfile(profileORM), nil
}
