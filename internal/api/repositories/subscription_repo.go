package repositories

import (
	"fmt"

	models "github.com/gene-qxsi/Flexive/internal/models/orm_models"
	"github.com/gene-qxsi/Flexive/internal/storage"
)

type SubscriptionRepo struct {
	storage *storage.Storage
}

func NewSubscriptionRepo(storage *storage.Storage) *SubscriptionRepo {
	return &SubscriptionRepo{
		storage: storage,
	}
}

func (r *SubscriptionRepo) CreateSubscription(subscription *models.Subscription) (int, int, error) {
	const op = "internal/api/repositories/subscription_repo.go/CreateSubscription()"

	err := r.storage.Sdb.Create(subscription).Error
	if err != nil {
		return 0, 0, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return subscription.UserID, subscription.Channel.ID, nil
}

func (r *SubscriptionRepo) GetSubscription(userID, channelID int) (*models.Subscription, error) {
	const op = "internal/api/repositories/subscription_repo.go/GetSubscription()"

	var subscription models.Subscription
	err := r.storage.Sdb.Preload("User").Preload("Channel").First(&subscription, userID, channelID).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &subscription, nil
}

func (r *SubscriptionRepo) GetSubscriptions() ([]models.Subscription, error) {
	const op = "internal/api/repositories/subscription_repo.go/GetSubscriptions()"

	var subscriptions []models.Subscription
	err := r.storage.Sdb.Preload("User").Preload("Channel").Find(&subscriptions).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return subscriptions, nil
}

func (r *SubscriptionRepo) DeleteSubscription(userID, channelID int) error {
	const op = "internal/api/repositories/subscription_repo.go/DeleteSubscription()"

	result := r.storage.Sdb.Delete(&models.Subscription{}, userID, channelID)
	if result.Error != nil {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-2: %s. ПУТЬ: %s",
			fmt.Sprintf("subscription with UserID %d AND channelID %d not found", userID, channelID), op)
	}

	return nil
}

func (r *SubscriptionRepo) UpdateSubscription(userID, channelID int, values map[string]interface{}) error {
	const op = "internal/api/repositories/subscription_repo.go/UpdateSubscription()"

	result := r.storage.Sdb.Model(&models.Subscription{}).Where("user_id = ? AND channel_id = ?", userID, channelID).Updates(values)
	if result.Error != nil {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s",
			fmt.Sprintf("subscription with UserID %d AND channelID %d not found or no changes made", userID, channelID), op)
	}

	return nil
}
