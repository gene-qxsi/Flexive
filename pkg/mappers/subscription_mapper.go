package mappers

import (
	"github.com/gene-qxsi/Flexive/internal/delivery/http/dto"
	"github.com/gene-qxsi/Flexive/internal/repository/models"
)

func SubcscriptionToDTO(orm models.Subscription) dto.SubscriptionDTO {
	return dto.SubscriptionDTO{
		UserID:    orm.UserID,
		ChannelID: orm.ChannelID,
		CreatedAt: orm.CreatedAt,
	}
}

func SubcscriptionsToDTOs(orm []models.Subscription) []dto.SubscriptionDTO {
	dto := make([]dto.SubscriptionDTO, len(orm))

	for i, subscription := range orm {
		dto[i] = SubcscriptionToDTO(subscription)
	}

	return dto
}
