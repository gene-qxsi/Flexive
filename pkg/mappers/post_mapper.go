package mappers

import (
	"github.com/gene-qxsi/Flexive/internal/delivery/http/dto"
	"github.com/gene-qxsi/Flexive/internal/repository/models"
)

func PostToDTO(orm models.Post) dto.PostDTO {
	return dto.PostDTO{
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

func PostsToDTOs(orm []models.Post) []dto.PostDTO {
	dto := make([]dto.PostDTO, len(orm))

	for i, post := range orm {
		dto[i] = PostToDTO(post)
	}

	return dto
}
