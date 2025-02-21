package mappers

import (
	"github.com/gene-qxsi/Flexive/internal/delivery/http/dto"
	"github.com/gene-qxsi/Flexive/internal/repository/models"
)

func UserToDTO(orm models.User) dto.UserDTO {
	return dto.UserDTO{
		ID:            orm.ID,
		Username:      orm.Username,
		Description:   orm.Description,
		Birthday:      orm.Birthday,
		CreatedAt:     orm.CreatedAt,
		Role:          orm.Role,
		Password:      orm.Password,
		Channels:      ChannelsToDTOs(orm.Channels),
		Comments:      CommentsToDTOs(orm.Comments),
		Posts:         PostsToDTOs(orm.Posts),
		Reactions:     ReactionsToDTOs(orm.Reactions),
		Subscriptions: SubcscriptionsToDTOs(orm.Subscriptions),
	}
}

func UsersToDTOs(orm []models.User) []dto.UserDTO {
	dto := make([]dto.UserDTO, len(orm))

	for i, user := range orm {
		dto[i] = UserToDTO(user)
	}

	return dto
}
