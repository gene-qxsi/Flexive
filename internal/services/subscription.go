package services

import (
	"github.com/gene-qxsi/Flexive/internal/dto"
	"github.com/gene-qxsi/Flexive/internal/repository/models"
	repository "github.com/gene-qxsi/Flexive/internal/repository/sqlrepo"
)

type SubscriptionService struct {
	Repo *repository.SubscriptionRepo
}

func NewSubscriptionService(repo *repository.SubscriptionRepo) *SubscriptionService {
	return &SubscriptionService{Repo: repo}
}

func (s *SubscriptionService) CreateSubscription(subscription *models.Subscription) (*dto.SubscriptionDTO, error) {
	// subscription, err := s.Repo.CreateSubscription(subscription)
	// if err != nil {
	// 	return nil, err
	// }

	// subscriptionDTO := mappers.SubcscriptionToDTO(*subscription)
	return nil, nil
}

func (s *SubscriptionService) GetSubscription(userID, channelID int) (*dto.SubscriptionDTO, error) {
	// subscription, err := s.Repo.GetSubscription(userID, channelID)
	// if err != nil {
	// 	return nil, err
	// }

	// subscriptionDTO := mappers.SubcscriptionToDTO(*subscription)
	return nil, nil
}

func (s *SubscriptionService) GetSubscriptions() ([]dto.SubscriptionDTO, error) {
	// subscriptions, err := s.Repo.GetSubscriptions()
	// if err != nil {
	// 	return nil, err
	// }

	// subscriptionsDTOs := mappers.SubcscriptionsToDTOs(subscriptions)
	return nil, nil
}

func (s *SubscriptionService) DeleteSubscription(userID, channelID int) error {
	return s.Repo.DeleteSubscription(userID, channelID)
}

func (s *SubscriptionService) UpdateSubscription(userID, channelID int, values map[string]interface{}) (*dto.SubscriptionDTO, error) {
	// subscription, err := s.Repo.UpdateSubscription(userID, channelID, values)
	// if err != nil {
	// 	return nil, err
	// }

	// subscriptionDTO := mappers.SubcscriptionToDTO(*subscription)
	return nil, nil
}
