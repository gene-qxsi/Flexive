package mappers

import (
	"github.com/gene-qxsi/Flexive/internal/models/dto_models"
	"github.com/gene-qxsi/Flexive/internal/models/orm_models"
)

func SubcscriptionToDTO(orm orm_models.Subscription) dto_models.SubscriptionDTO {
	return dto_models.SubscriptionDTO{
		UserID:    orm.UserID,
		ChannelID: orm.ChannelID,
		CreatedAt: orm.CreatedAt,
	}
}

func SubcscriptionsToDTOs(orm []orm_models.Subscription) []dto_models.SubscriptionDTO {
	dto := make([]dto_models.SubscriptionDTO, len(orm))

	for i, subscription := range orm {
		dto[i] = SubcscriptionToDTO(subscription)
	}

	return dto
}
