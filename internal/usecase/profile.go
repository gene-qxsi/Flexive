package usecase

import (
	"context"

	"github.com/gene-qxsi/Flexive/internal/delivery/http/v1/dto"
	"github.com/gene-qxsi/Flexive/internal/services"
)

// type ProfileUsecaseInterface interface {
// 	GetMyProfile(ctx context.Context) (*dto.Profile, error)
// 	// GetProfile(ctx context.Context, ) (domain.Profile, error)
// 	// UpdateProfile(ctx context.Context, UpdProfile dto.UpdateProfile) (domain.Profile, error)
// }

type ProfileUsecase struct {
	profileSrv *services.ProfileService
}

func NewProfileUsecase(profileSrv *services.ProfileService) *ProfileUsecase {
	return &ProfileUsecase{profileSrv: profileSrv}
}

func (u *ProfileUsecase) GetProfileByUserID(ctx context.Context, userID int) (*dto.Profile, error) {

	profile, err := u.profileSrv.GetProfileByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	profileDTO := dto.Profile{
		Bio:       profile.Bio,
		UserID:    profile.UserID,
		Username:  profile.Username,
		AvatarURL: profile.AvatarURL,
		Website:   profile.Website,
		Birthday:  profile.Birthday,
		Role:      profile.Role,
	}

	return &profileDTO, nil
}

func (u *ProfileUsecase) UpdateProfile(ctx context.Context, userID int, values map[string]interface{}) (*dto.Profile, error) {

	profile, err := u.profileSrv.UpdateProfileByUserID(ctx, userID, values)
	if err != nil {
		return nil, err
	}

	profileDTO := dto.Profile{
		UserID:    profile.UserID,
		Username:  profile.Username,
		Role:      profile.Role,
		Bio:       profile.Bio,
		AvatarURL: profile.AvatarURL,
		Website:   profile.Website,
		Birthday:  profile.Birthday,
	}

	return &profileDTO, nil
}

func (u *ProfileUsecase) GetProfiles(ctx context.Context) ([]dto.Profile, error) {

	profilesDomain, err := u.profileSrv.GetProfiles(ctx)
	if err != nil {
		return nil, err
	}

	var profiles []dto.Profile
	for _, profile := range profilesDomain {
		profiles = append(profiles, dto.Profile{
			UserID:    profile.UserID,
			Username:  profile.Username,
			Bio:       profile.Bio,
			AvatarURL: profile.AvatarURL,
			Website:   profile.Website,
			Birthday:  profile.Birthday,
			Role:      profile.Role,
		})
	}

	return profiles, nil
}
