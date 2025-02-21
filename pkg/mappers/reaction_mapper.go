package mappers

import (
	"github.com/gene-qxsi/Flexive/internal/delivery/http/dto"
	"github.com/gene-qxsi/Flexive/internal/repository/models"
)

func ReactionToDTO(orm models.Reaction) dto.ReactionDTO {
	return dto.ReactionDTO{
		UserID:    orm.UserID,
		PostID:    orm.PostID,
		Reaction:  orm.Reaction,
		CreatedAt: orm.CreatedAt,
	}
}

func ReactionsToDTOs(orm []models.Reaction) []dto.ReactionDTO {
	dto := make([]dto.ReactionDTO, len(orm))

	for i, reaction := range orm {
		dto[i] = ReactionToDTO(reaction)
	}

	return dto
}
