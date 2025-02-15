package services

import (
	"github.com/gene-qxsi/Flexive/internal/api/repositories"
	models "github.com/gene-qxsi/Flexive/internal/models/orm_models"
)

type PostService struct {
	Repo *repositories.PostRepo
}

func NewPostService(repo *repositories.PostRepo) *PostService {
	return &PostService{Repo: repo}
}

func (s *PostService) CreatePost(post *models.Post) (int, error) {
	return s.Repo.CreatePost(post)
}

func (s *PostService) GetPost(id int) (*models.Post, error) {
	return s.Repo.GetPost(id)
}

func (s *PostService) GetPosts() ([]models.Post, error) {
	return s.Repo.GetPosts()
}

func (s *PostService) DeletePost(id int) error {
	return s.Repo.DeletePost(id)
}

func (s *PostService) UpdatePost(id int, values map[string]interface{}) error {
	return s.Repo.UpdatePost(id, values)
}
