package repository

import (
	"fmt"

	"github.com/gene-qxsi/Flexive/internal/repository/models"
	"gorm.io/gorm"
)

type SubscriptionRepo struct {
	db *gorm.DB
}

func NewSubscriptionRepo(db *gorm.DB) *SubscriptionRepo {
	return &SubscriptionRepo{
		db: db,
	}
}

func (r *SubscriptionRepo) CreateSubscription(subscription *models.Subscription) (*models.Subscription, error) {
	const op = "internal/api/repositories/subscription_repo.go/CreateSubscription()"

	err := r.db.Create(subscription).Error
	if err != nil {
		return subscription, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return subscription, nil
}

func (r *SubscriptionRepo) GetSubscription(userID, channelID int) (*models.Subscription, error) {
	const op = "internal/api/repositories/subscription_repo.go/GetSubscription()"

	var subscription models.Subscription
	err := r.db.First(&subscription, userID, channelID).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &subscription, nil
}

func (r *SubscriptionRepo) GetSubscriptions() ([]models.Subscription, error) {
	const op = "internal/api/repositories/subscription_repo.go/GetSubscriptions()"

	var subscriptions []models.Subscription
	err := r.db.Find(&subscriptions).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return subscriptions, nil
}

func (r *SubscriptionRepo) DeleteSubscription(userID, channelID int) error {
	const op = "internal/api/repositories/subscription_repo.go/DeleteSubscription()"

	result := r.db.Delete(&models.Subscription{}, userID, channelID)
	if result.Error != nil {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-2: %s. ПУТЬ: %s",
			fmt.Sprintf("subscription with UserID %d AND channelID %d not found", userID, channelID), op)
	}

	return nil
}

func (r *SubscriptionRepo) UpdateSubscription(userID, channelID int, values map[string]interface{}) (*models.Subscription, error) {
	const op = "internal/api/repositories/subscription_repo.go/UpdateSubscription()"

	result := r.db.Model(&models.Subscription{}).Where("user_id = ? AND channel_id = ?", userID, channelID).Updates(values)
	if result.Error != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-2: %s. ПУТЬ: %s",
			fmt.Sprintf("subscription with UserID %d AND channelID %d not found or no changes made", userID, channelID), op)
	}

	var subscription *models.Subscription
	err := r.db.Where("user_id = ?", userID).Where("channel_id = ?", channelID).First(&subscription).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op)
	}

	return subscription, nil
}
