package mappers

import (
	"github.com/gene-qxsi/Flexive/internal/models/dto_models"
	"github.com/gene-qxsi/Flexive/internal/models/orm_models"
)

func ReactionToDTO(orm orm_models.Reaction) dto_models.ReactionDTO {
	return dto_models.ReactionDTO{
		UserID:    orm.UserID,
		PostID:    orm.PostID,
		Reaction:  orm.Reaction,
		CreatedAt: orm.CreatedAt,
	}
}

func ReactionsToDTOs(orm []orm_models.Reaction) []dto_models.ReactionDTO {
	dto := make([]dto_models.ReactionDTO, len(orm))

	for i, reaction := range orm {
		dto[i] = ReactionToDTO(reaction)
	}

	return dto
}
