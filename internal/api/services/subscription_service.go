package services

import (
	"github.com/gene-qxsi/Flexive/internal/api/repositories"
	"github.com/gene-qxsi/Flexive/internal/models/dto_models"
	models "github.com/gene-qxsi/Flexive/internal/models/orm_models"
	"github.com/gene-qxsi/Flexive/pkg/mappers"
)

type SubscriptionService struct {
	Repo *repositories.SubscriptionRepo
}

func NewSubscriptionService(repo *repositories.SubscriptionRepo) *SubscriptionService {
	return &SubscriptionService{Repo: repo}
}

func (s *SubscriptionService) CreateSubscription(subscription *models.Subscription) (*dto_models.SubscriptionDTO, error) {
	subscription, err := s.Repo.CreateSubscription(subscription)
	if err != nil {
		return nil, err
	}

	subscriptionDTO := mappers.SubcscriptionToDTO(*subscription)
	return &subscriptionDTO, err
}

func (s *SubscriptionService) GetSubscription(userID, channelID int) (*dto_models.SubscriptionDTO, error) {
	subscription, err := s.Repo.GetSubscription(userID, channelID)
	if err != nil {
		return nil, err
	}

	subscriptionDTO := mappers.SubcscriptionToDTO(*subscription)
	return &subscriptionDTO, err
}

func (s *SubscriptionService) GetSubscriptions() ([]dto_models.SubscriptionDTO, error) {
	subscriptions, err := s.Repo.GetSubscriptions()
	if err != nil {
		return nil, err
	}

	subscriptionsDTOs := mappers.SubcscriptionsToDTOs(subscriptions)
	return subscriptionsDTOs, err
}

func (s *SubscriptionService) DeleteSubscription(userID, channelID int) error {
	return s.Repo.DeleteSubscription(userID, channelID)
}

func (s *SubscriptionService) UpdateSubscription(userID, channelID int, values map[string]interface{}) (*dto_models.SubscriptionDTO, error) {
	subscription, err := s.Repo.UpdateSubscription(userID, channelID, values)
	if err != nil {
		return nil, err
	}

	subscriptionDTO := mappers.SubcscriptionToDTO(*subscription)
	return &subscriptionDTO, err
}
