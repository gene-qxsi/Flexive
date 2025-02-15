package services

import (
	"github.com/gene-qxsi/Flexive/internal/api/repositories"
	models "github.com/gene-qxsi/Flexive/internal/models/orm_models"
)

type ReactionService struct {
	Repo *repositories.ReactionRepo
}

func NewReactionService(repo *repositories.ReactionRepo) *ReactionService {
	return &ReactionService{Repo: repo}
}

func (s *ReactionService) CreateReaction(reaction *models.Reaction) (int, int, error) {
	return s.Repo.CreateReaction(reaction)
}

func (s *ReactionService) GetReaction(userID, postID int) (*models.Reaction, error) {
	return s.Repo.GetReaction(userID, postID)
}

func (s *ReactionService) GetReactions() ([]models.Reaction, error) {
	return s.Repo.GetReactions()
}

func (s *ReactionService) DeleteReaction(userID, postID int) error {
	return s.Repo.DeleteReaction(userID, postID)
}

func (s *ReactionService) UpdateReaction(userID, postID int, values map[string]interface{}) error {
	return s.Repo.UpdateReaction(userID, postID, values)
}
