package services

import (
	"github.com/gene-qxsi/Flexive/internal/api/repositories"
	models "github.com/gene-qxsi/Flexive/internal/models/orm_models"
)

type CommentService struct {
	Repo *repositories.CommentRepo
}

func NewCommentService(repo *repositories.CommentRepo) *CommentService {
	return &CommentService{Repo: repo}
}

func (s *CommentService) CreateComment(comment *models.Comment) (int, error) {
	return s.Repo.CreateComment(comment)
}

func (s *CommentService) GetComment(id int) (*models.Comment, error) {
	return s.Repo.GetComment(id)
}

func (s *CommentService) GetComments() ([]models.Comment, error) {
	return s.Repo.GetComments()
}

func (s *CommentService) DeleteComment(id int) error {
	return s.Repo.DeleteComment(id)
}

func (s *CommentService) UpdateComment(id int, values map[string]interface{}) error {
	return s.Repo.UpdateComment(id, values)
}
