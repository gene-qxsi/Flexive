package usecase

import (
	"context"

	"github.com/gene-qxsi/Flexive/internal/delivery/http/dto"
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
		AvatarURL: profile.AvatarURL,
		Website:   profile.Website,
		Birthday:  profile.Birthday}

	return &profileDTO, nil
}

func (u *ProfileUsecase) UpdateProfile(ctx context.Context, userID int, values map[string]string) (*dto.Profile, error) {

	profile, err := u.profileSrv.UpdateProfileByUserID(ctx, userID, values)
	if err != nil {
		return nil, err
	}

	profileDTO := dto.Profile{
		Bio:       profile.Bio,
		AvatarURL: profile.AvatarURL,
		Website:   profile.Website,
		Birthday:  profile.Birthday,
	}

	return &profileDTO, nil
}
