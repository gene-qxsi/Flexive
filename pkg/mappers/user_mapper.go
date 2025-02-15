package mappers

import (
	"github.com/gene-qxsi/Flexive/internal/models/dto_models"
	"github.com/gene-qxsi/Flexive/internal/models/orm_models"
)

func UserToDTO(orm orm_models.User) dto_models.UserDTO {
	return dto_models.UserDTO{
		ID:            orm.ID,
		Username:      orm.Username,
		Description:   orm.Description,
		Birthday:      orm.Birthday,
		CreatedAt:     orm.CreatedAt,
		Channels:      ChannelsToDTOs(orm.Channels),
		Comments:      CommentsToDTOs(orm.Comments),
		Posts:         PostsToDTOs(orm.Posts),
		Reactions:     ReactionsToDTOs(orm.Reactions),
		Subscriptions: SubcscriptionsToDTOs(orm.Subscriptions),
	}
}

func UsersToDTOs(orm []orm_models.User) []dto_models.UserDTO {
	dto := make([]dto_models.UserDTO, len(orm))

	for i, user := range orm {
		dto[i] = UserToDTO(user)
	}

	return dto
}
