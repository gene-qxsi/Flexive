package services

import (
	"github.com/gene-qxsi/Flexive/internal/api/repositories"
	models "github.com/gene-qxsi/Flexive/internal/models/orm_models"
)

type SubscriptionService struct {
	Repo *repositories.SubscriptionRepo
}

func NewSubscriptionService(repo *repositories.SubscriptionRepo) *SubscriptionService {
	return &SubscriptionService{Repo: repo}
}

func (s *SubscriptionService) CreateSubscription(subscription *models.Subscription) (int, int, error) {
	return s.Repo.CreateSubscription(subscription)
}

func (s *SubscriptionService) GetSubscription(userID, channelID int) (*models.Subscription, error) {
	return s.Repo.GetSubscription(userID, channelID)
}

func (s *SubscriptionService) GetSubscriptions() ([]models.Subscription, error) {
	return s.Repo.GetSubscriptions()
}

func (s *SubscriptionService) DeleteSubscription(userID, channelID int) error {
	return s.Repo.DeleteSubscription(userID, channelID)
}

func (s *SubscriptionService) UpdateSubscription(userID, channelID int, values map[string]interface{}) error {
	return s.Repo.UpdateSubscription(userID, channelID, values)
}
