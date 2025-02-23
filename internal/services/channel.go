package services

import (
	"time"

	"github.com/gene-qxsi/Flexive/internal/delivery/http/dto"
	"github.com/gene-qxsi/Flexive/internal/repository"
	"github.com/gene-qxsi/Flexive/internal/repository/models"
	"github.com/gene-qxsi/Flexive/pkg/mappers"
)

type Channel struct {
	ID          int         `gorm:"type:int;primaryKey" json:"id"`
	Title       string      `gorm:"type:varchar(64);size:64;not null" json:"title"`
	Description string      `gorm:"type:varchar(1024);size:1024" json:"description"`
	UserID      int         `gorm:"type:bigint;not null;index" json:"user_id"`
	User        models.User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
	CreatedAt   time.Time   `gorm:"type:timestamp;default:now()" json:"created_at"`
}

type ChannelService struct {
	Repo *repository.ChannelRepo
}

func NewChannelService(repo *repository.ChannelRepo) *ChannelService {
	return &ChannelService{Repo: repo}
}

func (s *ChannelService) CreateChannel(channel *models.Channel) (*dto.ChannelDTO, error) {
	channel, err := s.Repo.CreateChannel(channel)
	if err != nil {
		return nil, err
	}

	channelDTO := mappers.ChannelToDTO(*channel)
	return &channelDTO, nil
}

func (s *ChannelService) GetChannel(id int) (*dto.ChannelDTO, error) {
	channel, err := s.Repo.GetChannel(id)
	if err != nil {
		return nil, err
	}

	channelDTO := mappers.ChannelToDTO(*channel)
	return &channelDTO, nil
}

func (s *ChannelService) GetChannels() ([]dto.ChannelDTO, error) {
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

func (s *ChannelService) UpdateChannel(id int, values map[string]interface{}) (*dto.ChannelDTO, error) {
	channel, err := s.Repo.UpdateChannel(id, values)
	if err != nil {
		return nil, err
	}

	channelDTO := mappers.ChannelToDTO(*channel)
	return &channelDTO, nil
}
