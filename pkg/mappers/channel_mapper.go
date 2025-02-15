package mappers

import (
	"github.com/gene-qxsi/Flexive/internal/models/dto_models"
	"github.com/gene-qxsi/Flexive/internal/models/orm_models"
)

func ChannelToDTO(orm orm_models.Channel) dto_models.ChannelDTO {
	return dto_models.ChannelDTO{
		ID:          orm.ID,
		Title:       orm.Title,
		Description: orm.Description,
		UserID:      orm.UserID,
		CreatedAt:   orm.CreatedAt,
	}
}
func ChannelsToDTOs(orm []orm_models.Channel) []dto_models.ChannelDTO {
	dto := make([]dto_models.ChannelDTO, len(orm))

	for i, channel := range orm {
		dto[i] = ChannelToDTO(channel)
	}

	return dto
}
