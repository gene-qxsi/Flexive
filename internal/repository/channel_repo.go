package repository

import (
	"fmt"

	"github.com/gene-qxsi/Flexive/internal/repository/models"
	"gorm.io/gorm"
)

type ChannelRepo struct {
	db *gorm.DB
}

func NewChannelRepo(db *gorm.DB) *ChannelRepo {
	return &ChannelRepo{
		db: db,
	}
}

func (r *ChannelRepo) CreateChannel(channel *models.Channel) (*models.Channel, error) {
	const op = "internal/api/repositories/channel_repo.go/CreateChannel()"

	err := r.db.Create(channel).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return channel, nil
}

func (r *ChannelRepo) GetChannel(id int) (*models.Channel, error) {
	const op = "internal/api/repositories/channel_repo.go/GetChannel()"

	var channel models.Channel
	err := r.db.First(&channel, id).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &channel, nil
}

func (r *ChannelRepo) GetChannels() ([]models.Channel, error) {
	const op = "internal/api/repositories/channel_repo.go/GetChannels()"

	var channels []models.Channel
	err := r.db.Find(&channels).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return channels, nil
}

func (r *ChannelRepo) DeleteChannel(id int) error {
	const op = "internal/api/repositories/channel_repo.go/DeleteChannel()"

	result := r.db.Delete(&models.Channel{}, id)
	if result.Error != nil {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-2: %s. ПУТЬ: %s", fmt.Sprintf("channel with ID %d not found", id), op)
	}

	return nil
}

func (r *ChannelRepo) UpdateChannel(id int, values map[string]interface{}) (*models.Channel, error) {
	const op = "internal/api/repositories/channel_repo.go/UpdateChannel()"

	result := r.db.Model(&models.Channel{}).Where("id = ?", id).Updates(values)
	if result.Error != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", fmt.Sprintf("channel with ID %d not found or no changes made", id), op)
	}

	var channel models.Channel
	err := r.db.Where("id = ?", id).First(&channel).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &channel, nil
}
