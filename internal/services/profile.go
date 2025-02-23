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

func (s *ProfileService) GetProfileByUserID(ctx context.Context, userID int) (*domain.Profile, error) {

	profileORM, err := s.repo.GetProfileByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return domain.ToDomainProfile(profileORM), nil
}

func (s *ProfileService) UpdateProfileByUserID(ctx context.Context, userID int, values map[string]string) (*domain.Profile, error) {

	profileORM, err := s.repo.UpdateProfileByUserID(ctx, userID, values)
	if err != nil {
		return nil, err
	}

	return domain.ToDomainProfile(profileORM), nil
}
