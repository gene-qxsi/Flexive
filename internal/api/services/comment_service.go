package services

import (
	"github.com/gene-qxsi/Flexive/internal/api/repositories"
	"github.com/gene-qxsi/Flexive/internal/models/dto_models"
	models "github.com/gene-qxsi/Flexive/internal/models/orm_models"
	"github.com/gene-qxsi/Flexive/pkg/mappers"
)

type CommentService struct {
	Repo *repositories.CommentRepo
}

func NewCommentService(repo *repositories.CommentRepo) *CommentService {
	return &CommentService{Repo: repo}
}

func (s *CommentService) CreateComment(comment *models.Comment) (*dto_models.CommentDTO, error) {
	comment, err := s.Repo.CreateComment(comment)
	if err != nil {
		return nil, err
	}

	commentDTO := mappers.CommentToDTO(*comment)
	return &commentDTO, nil
}

func (s *CommentService) GetComment(id int) (*dto_models.CommentDTO, error) {
	comment, err := s.Repo.GetComment(id)
	if err != nil {
		return nil, err
	}

	commentDTO := mappers.CommentToDTO(*comment)
	return &commentDTO, nil
}

func (s *CommentService) GetComments() ([]dto_models.CommentDTO, error) {
	comments, err := s.Repo.GetComments()
	if err != nil {
		return nil, err
	}

	commentsDTOs := mappers.CommentsToDTOs(comments)
	return commentsDTOs, nil
}

func (s *CommentService) DeleteComment(id int) error {
	return s.Repo.DeleteComment(id)
}

func (s *CommentService) UpdateComment(id int, values map[string]interface{}) (*dto_models.CommentDTO, error) {
	comment, err := s.Repo.UpdateComment(id, values)
	if err != nil {
		return nil, err
	}

	commentDTO := mappers.CommentToDTO(*comment)
	return &commentDTO, nil
}
