package services

import (
	"github.com/gene-qxsi/Flexive/internal/api/repositories"
	models "github.com/gene-qxsi/Flexive/internal/models/orm_models"
)

type ChannelService struct {
	Repo *repositories.ChannelRepo
}

func NewChannelService(repo *repositories.ChannelRepo) *ChannelService {
	return &ChannelService{Repo: repo}
}

func (s *ChannelService) CreateChannel(channel *models.Channel) (int, error) {
	return s.Repo.CreateChannel(channel)
}

func (s *ChannelService) GetChannel(id int) (*models.Channel, error) {
	return s.Repo.GetChannel(id)
}

func (s *ChannelService) GetChannels() ([]models.Channel, error) {
	return s.Repo.GetChannels()
}

func (s *ChannelService) DeleteChannel(id int) error {
	return s.Repo.DeleteChannel(id)
}

func (s *ChannelService) UpdateChannel(id int, values map[string]interface{}) error {
	return s.Repo.UpdateChannel(id, values)
}
