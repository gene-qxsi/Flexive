package services

import (
	"github.com/gene-qxsi/Flexive/internal/api/repositories"
	"github.com/gene-qxsi/Flexive/internal/models/dto_models"
	models "github.com/gene-qxsi/Flexive/internal/models/orm_models"
	"github.com/gene-qxsi/Flexive/pkg/mappers"
)

type PostService struct {
	Repo *repositories.PostRepo
}

func NewPostService(repo *repositories.PostRepo) *PostService {
	return &PostService{Repo: repo}
}

func (s *PostService) CreatePost(post *models.Post) (*dto_models.PostDTO, error) {
	post, err := s.Repo.CreatePost(post)
	if err != nil {
		return nil, err
	}

	postDTO := mappers.PostToDTO(*post)
	return &postDTO, nil
}

func (s *PostService) GetPost(id int) (*dto_models.PostDTO, error) {
	post, err := s.Repo.GetPost(id)
	if err != nil {
		return nil, err
	}

	postDTO := mappers.PostToDTO(*post)
	return &postDTO, nil
}

func (s *PostService) GetPosts() ([]dto_models.PostDTO, error) {
	posts, err := s.Repo.GetPosts()
	if err != nil {
		return nil, err
	}

	postsDTOs := mappers.PostsToDTOs(posts)
	return postsDTOs, nil
}

func (s *PostService) DeletePost(id int) error {
	return s.Repo.DeletePost(id)
}

func (s *PostService) UpdatePost(id int, values map[string]interface{}) (*dto_models.PostDTO, error) {
	post, err := s.Repo.UpdatePost(id, values)
	if err != nil {
		return nil, err
	}

	postDTO := mappers.PostToDTO(*post)
	return &postDTO, nil
}
