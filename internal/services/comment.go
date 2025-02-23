package services

import (
	"github.com/gene-qxsi/Flexive/internal/delivery/http/v1/dto"
	"github.com/gene-qxsi/Flexive/internal/repository"
	"github.com/gene-qxsi/Flexive/internal/repository/models"
)

type CommentService struct {
	Repo *repository.CommentRepo
}

func NewCommentService(repo *repository.CommentRepo) *CommentService {
	return &CommentService{Repo: repo}
}

func (s *CommentService) CreateComment(comment *models.Comment) (*dto.CommentDTO, error) {
	// comment, err := s.Repo.CreateComment(comment)
	// if err != nil {
	// 	return nil, err
	// }

	// commentDTO := mappers.CommentToDTO(*comment)
	return nil, nil
}

func (s *CommentService) GetComment(id int) (*dto.CommentDTO, error) {
	// comment, err := s.Repo.GetComment(id)
	// if err != nil {
	// 	return nil, err
	// }

	// commentDTO := mappers.CommentToDTO(*comment)
	return nil, nil
}

func (s *CommentService) GetComments() ([]dto.CommentDTO, error) {
	// comments, err := s.Repo.GetComments()
	// if err != nil {
	// 	return nil, err
	// }

	// commentsDTOs := mappers.CommentsToDTOs(comments)
	return nil, nil
}

func (s *CommentService) DeleteComment(id int) error {
	return s.Repo.DeleteComment(id)
}

func (s *CommentService) UpdateComment(id int, values map[string]interface{}) (*dto.CommentDTO, error) {
	// comment, err := s.Repo.UpdateComment(id, values)
	// if err != nil {
	// 	return nil, err
	// }

	// commentDTO := mappers.CommentToDTO(*comment)
	return nil, nil
}
