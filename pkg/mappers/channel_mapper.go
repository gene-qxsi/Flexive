package mappers

import (
	"github.com/gene-qxsi/Flexive/internal/delivery/http/dto"
	"github.com/gene-qxsi/Flexive/internal/repository/models"
)

func ChannelToDTO(orm models.Channel) dto.ChannelDTO {
	return dto.ChannelDTO{
		ID:          orm.ID,
		Title:       orm.Title,
		Description: orm.Description,
		UserID:      orm.UserID,
		CreatedAt:   orm.CreatedAt,
	}
}
func ChannelsToDTOs(orm []models.Channel) []dto.ChannelDTO {
	dto := make([]dto.ChannelDTO, len(orm))

	for i, channel := range orm {
		dto[i] = ChannelToDTO(channel)
	}

	return dto
}
