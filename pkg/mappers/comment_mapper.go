package mappers

import (
	"github.com/gene-qxsi/Flexive/internal/delivery/http/dto"
	"github.com/gene-qxsi/Flexive/internal/repository/models"
)

func CommentToDTO(orm models.Comment) dto.CommentDTO {
	return dto.CommentDTO{
		ID:        orm.ID,
		Content:   orm.Content,
		UserID:    orm.UserID,
		PostID:    orm.PostID,
		CreatedAt: orm.CreatedAt,
	}
}

func CommentsToDTOs(orm []models.Comment) []dto.CommentDTO {
	dto := make([]dto.CommentDTO, len(orm))

	for i, comment := range orm {
		dto[i] = CommentToDTO(comment)
	}

	return dto
}
