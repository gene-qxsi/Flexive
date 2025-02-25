package services

import (
	"github.com/gene-qxsi/Flexive/internal/delivery/http/v1/dto"
	"github.com/gene-qxsi/Flexive/internal/repository/models"
	repository "github.com/gene-qxsi/Flexive/internal/repository/sqlrepo"
)

type PostService struct {
	Repo *repository.PostRepo
}

func NewPostService(repo *repository.PostRepo) *PostService {
	return &PostService{Repo: repo}
}

func (s *PostService) CreatePost(post *models.Post) (*dto.PostDTO, error) {
	// post, err := s.Repo.CreatePost(post)
	// if err != nil {
	// 	return nil, err
	// }

	// postDTO := mappers.PostToDTO(*post)
	return nil, nil
}

func (s *PostService) GetPost(id int) (*dto.PostDTO, error) {
	// post, err := s.Repo.GetPost(id)
	// if err != nil {
	// 	return nil, err
	// }

	// postDTO := mappers.PostToDTO(*post)
	return nil, nil
}

func (s *PostService) GetPosts() ([]dto.PostDTO, error) {
	// posts, err := s.Repo.GetPosts()
	// if err != nil {
	// 	return nil, err
	// }

	// postsDTOs := mappers.PostsToDTOs(posts)
	return nil, nil
}

func (s *PostService) DeletePost(id int) error {
	return s.Repo.DeletePost(id)
}

func (s *PostService) UpdatePost(id int, values map[string]interface{}) (*dto.PostDTO, error) {
	// post, err := s.Repo.UpdatePost(id, values)
	// if err != nil {
	// 	return nil, err
	// }

	// postDTO := mappers.PostToDTO(*post)
	return nil, nil
}
