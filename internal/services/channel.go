package services

import (
	"github.com/gene-qxsi/Flexive/internal/domain"
	"github.com/gene-qxsi/Flexive/internal/repository/models"
	repository "github.com/gene-qxsi/Flexive/internal/repository/sqlrepo"
)

type ChannelService struct {
	Repo *repository.ChannelRepo
}

func NewChannelService(repo *repository.ChannelRepo) *ChannelService {
	return &ChannelService{Repo: repo}
}

func (s *ChannelService) CreateChannel(channelORM *models.Channel) (*domain.Channel, error) {
	channelORM, err := s.Repo.CreateChannel(channelORM)
	if err != nil {
		return nil, err
	}

	channel := domain.ToDomainChannel(channelORM)
	return channel, nil
}

func (s *ChannelService) GetChannel(id int) (*domain.Channel, error) {
	channel, err := s.Repo.GetChannel(id)
	if err != nil {
		return nil, err
	}

	channelDTO := domain.ToDomainChannel(channel)
	return channelDTO, nil
}

func (s *ChannelService) GetChannels() ([]domain.Channel, error) {
	channelsORMs, err := s.Repo.GetChannels()
	if err != nil {
		return nil, err
	}
	var channels []domain.Channel
	for _, channel := range channelsORMs {
		channels = append(channels, *domain.ToDomainChannel(&channel))
	}

	return channels, nil
}

func (s *ChannelService) DeleteChannel(id int) error {
	return s.Repo.DeleteChannel(id)
}

func (s *ChannelService) UpdateChannel(id int, values map[string]interface{}) (*domain.Channel, error) {
	channel, err := s.Repo.UpdateChannel(id, values)
	if err != nil {
		return nil, err
	}

	channelDTO := domain.ToDomainChannel(channel)
	return channelDTO, nil
}
