package mappers

import (
	"github.com/gene-qxsi/Flexive/internal/models/dto_models"
	"github.com/gene-qxsi/Flexive/internal/models/orm_models"
)

func CommentToDTO(orm orm_models.Comment) dto_models.CommentDTO {
	return dto_models.CommentDTO{
		ID:        orm.ID,
		Content:   orm.Content,
		UserID:    orm.UserID,
		PostID:    orm.PostID,
		CreatedAt: orm.CreatedAt,
	}
}

func CommentsToDTOs(orm []orm_models.Comment) []dto_models.CommentDTO {
	dto := make([]dto_models.CommentDTO, len(orm))

	for i, comment := range orm {
		dto[i] = CommentToDTO(comment)
	}

	return dto
}
