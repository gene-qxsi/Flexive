package services

import (
	"github.com/gene-qxsi/Flexive/internal/api/repositories"
	"github.com/gene-qxsi/Flexive/internal/models/dto_models"
	models "github.com/gene-qxsi/Flexive/internal/models/orm_models"
	"github.com/gene-qxsi/Flexive/pkg/mappers"
)

type ChannelService struct {
	Repo *repositories.ChannelRepo
}

func NewChannelService(repo *repositories.ChannelRepo) *ChannelService {
	return &ChannelService{Repo: repo}
}

func (s *ChannelService) CreateChannel(channel *models.Channel) (*dto_models.ChannelDTO, error) {
	channel, err := s.Repo.CreateChannel(channel)
	if err != nil {
		return nil, err
	}

	channelDTO := mappers.ChannelToDTO(*channel)
	return &channelDTO, nil
}

func (s *ChannelService) GetChannel(id int) (*dto_models.ChannelDTO, error) {
	channel, err := s.Repo.GetChannel(id)
	if err != nil {
		return nil, err
	}

	channelDTO := mappers.ChannelToDTO(*channel)
	return &channelDTO, nil
}

func (s *ChannelService) GetChannels() ([]dto_models.ChannelDTO, error) {
	channels, err := s.Repo.GetChannels()
	if err != nil {
		return nil, err
	}

	channelsDTOs := mappers.ChannelsToDTOs(channels)
	return channelsDTOs, nil
}

func (s *ChannelService) DeleteChannel(id int) error {
	return s.Repo.DeleteChannel(id)
}

func (s *ChannelService) UpdateChannel(id int, values map[string]interface{}) (*dto_models.ChannelDTO, error) {
	channel, err := s.Repo.UpdateChannel(id, values)
	if err != nil {
		return nil, err
	}

	channelDTO := mappers.ChannelToDTO(*channel)
	return &channelDTO, nil
}
