package mappers

import (
	"github.com/gene-qxsi/Flexive/internal/models/dto_models"
	"github.com/gene-qxsi/Flexive/internal/models/orm_models"
)

func PostToDTO(orm orm_models.Post) dto_models.PostDTO {
	return dto_models.PostDTO{
		ID:          orm.ID,
		Title:       orm.Title,
		Description: orm.Description,
		Content:     orm.Content,
		Views:       orm.Views,
		UserID:      orm.UserID,
		ChannelID:   orm.ChannelID,
		CreatedAt:   orm.CreatedAt,
	}
}

func PostsToDTOs(orm []orm_models.Post) []dto_models.PostDTO {
	dto := make([]dto_models.PostDTO, len(orm))

	for i, post := range orm {
		dto[i] = PostToDTO(post)
	}

	return dto
}
