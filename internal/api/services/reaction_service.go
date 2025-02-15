package services

import (
	"github.com/gene-qxsi/Flexive/internal/api/repositories"
	"github.com/gene-qxsi/Flexive/internal/models/dto_models"
	models "github.com/gene-qxsi/Flexive/internal/models/orm_models"
	"github.com/gene-qxsi/Flexive/pkg/mappers"
)

type ReactionService struct {
	Repo *repositories.ReactionRepo
}

func NewReactionService(repo *repositories.ReactionRepo) *ReactionService {
	return &ReactionService{Repo: repo}
}

func (s *ReactionService) CreateReaction(reaction *models.Reaction) (*dto_models.ReactionDTO, error) {
	reaction, err := s.Repo.CreateReaction(reaction)
	if err != nil {
		return nil, err
	}

	reactionDTO := mappers.ReactionToDTO(*reaction)
	return &reactionDTO, nil
}

func (s *ReactionService) GetReaction(userID, postID int) (*dto_models.ReactionDTO, error) {
	reactions, err := s.Repo.GetReaction(userID, postID)
	if err != nil {
		return nil, err
	}

	reactionDTO := mappers.ReactionToDTO(*reactions)
	return &reactionDTO, nil
}

func (s *ReactionService) GetReactions() ([]dto_models.ReactionDTO, error) {
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

func (s *ReactionService) UpdateReaction(userID, postID int, values map[string]interface{}) (*dto_models.ReactionDTO, error) {
	reactiuon, err := s.Repo.UpdateReaction(userID, postID, values)
	if err != nil {
		return nil, err
	}

	reactionDTO := mappers.ReactionToDTO(*reactiuon)
	return &reactionDTO, nil
}
