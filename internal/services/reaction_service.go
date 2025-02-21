package services

import (
	"github.com/gene-qxsi/Flexive/internal/delivery/http/dto"
	"github.com/gene-qxsi/Flexive/internal/repository"
	"github.com/gene-qxsi/Flexive/internal/repository/models"
	"github.com/gene-qxsi/Flexive/pkg/mappers"
)

type ReactionService struct {
	Repo *repository.ReactionRepo
}

func NewReactionService(repo *repository.ReactionRepo) *ReactionService {
	return &ReactionService{Repo: repo}
}

func (s *ReactionService) CreateReaction(reaction *models.Reaction) (*dto.ReactionDTO, error) {
	reaction, err := s.Repo.CreateReaction(reaction)
	if err != nil {
		return nil, err
	}

	reactionDTO := mappers.ReactionToDTO(*reaction)
	return &reactionDTO, nil
}

func (s *ReactionService) GetReaction(userID, postID int) (*dto.ReactionDTO, error) {
	reactions, err := s.Repo.GetReaction(userID, postID)
	if err != nil {
		return nil, err
	}

	reactionDTO := mappers.ReactionToDTO(*reactions)
	return &reactionDTO, nil
}

func (s *ReactionService) GetReactions() ([]dto.ReactionDTO, error) {
	reactions, err := s.Repo.GetReactions()
	if err != nil {
		return nil, err
	}
	reactionsDTOs := mappers.ReactionsToDTOs(reactions)
	return reactionsDTOs, nil
}

func (s *ReactionService) DeleteReaction(userID, postID int) error {
	return s.Repo.DeleteReaction(userID, postID)
}

func (s *ReactionService) UpdateReaction(userID, postID int, values map[string]interface{}) (*dto.ReactionDTO, error) {
	reactiuon, err := s.Repo.UpdateReaction(userID, postID, values)
	if err != nil {
		return nil, err
	}

	reactionDTO := mappers.ReactionToDTO(*reactiuon)
	return &reactionDTO, nil
}
