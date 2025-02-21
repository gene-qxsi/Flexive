package repository

import (
	"fmt"

	"github.com/gene-qxsi/Flexive/internal/repository/models"
	"github.com/gene-qxsi/Flexive/internal/storage"
)

type ChannelRepo struct {
	storage *storage.Storage
}

func NewChannelRepo(storage *storage.Storage) *ChannelRepo {
	return &ChannelRepo{
		storage: storage,
	}
}

func (r *ChannelRepo) CreateChannel(channel *models.Channel) (*models.Channel, error) {
	const op = "internal/api/repositories/channel_repo.go/CreateChannel()"

	err := r.storage.Sdb.Create(channel).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return channel, nil
}

func (r *ChannelRepo) GetChannel(id int) (*models.Channel, error) {
	const op = "internal/api/repositories/channel_repo.go/GetChannel()"

	var channel models.Channel
	err := r.storage.Sdb.First(&channel, id).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &channel, nil
}

func (r *ChannelRepo) GetChannels() ([]models.Channel, error) {
	const op = "internal/api/repositories/channel_repo.go/GetChannels()"

	var channels []models.Channel
	err := r.storage.Sdb.Find(&channels).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return channels, nil
}

func (r *ChannelRepo) DeleteChannel(id int) error {
	const op = "internal/api/repositories/channel_repo.go/DeleteChannel()"

	result := r.storage.Sdb.Delete(&models.Channel{}, id)
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

	result := r.storage.Sdb.Model(&models.Channel{}).Where("id = ?", id).Updates(values)
	if result.Error != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", fmt.Sprintf("channel with ID %d not found or no changes made", id), op)
	}

	var channel models.Channel
	err := r.storage.Sdb.Where("id = ?", id).First(&channel).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &channel, nil
}
