package repository

import (
	"fmt"

	"github.com/gene-qxsi/Flexive/internal/repository/models"
	"github.com/gene-qxsi/Flexive/internal/storage"
)

type PostRepo struct {
	storage *storage.Storage
}

func NewPostRepo(storage *storage.Storage) *PostRepo {
	return &PostRepo{
		storage: storage,
	}
}

func (r *PostRepo) CreatePost(post *models.Post) (*models.Post, error) {
	const op = "internal/api/repositories/post_repo.go/CreatePost()"

	err := r.storage.Sdb.Create(post).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return post, nil
}

func (r *PostRepo) GetPost(id int) (*models.Post, error) {
	const op = "internal/api/repositories/post_repo.go/GetPost()"

	var post models.Post
	err := r.storage.Sdb.First(&post, id).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &post, nil
}

func (r *PostRepo) GetPosts() ([]models.Post, error) {
	const op = "internal/api/repositories/post_repo.go/GetPosts()"

	var posts []models.Post
	err := r.storage.Sdb.Find(&posts).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return posts, nil
}

func (r *PostRepo) DeletePost(id int) error {
	const op = "internal/api/repositories/post_repo.go/DeletePost()"

	result := r.storage.Sdb.Delete(&models.Post{}, id)
	if result.Error != nil {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-2: %s. ПУТЬ: %s", fmt.Sprintf("post with ID %d not found", id), op)
	}

	return nil
}

func (r *PostRepo) UpdatePost(id int, values map[string]interface{}) (*models.Post, error) {
	const op = "internal/api/repositories/post_repo.go/UpdatePost()"

	result := r.storage.Sdb.Model(&models.Post{}).Where("id = ?", id).Updates(values)
	if result.Error != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", fmt.Sprintf("post with ID %d not found or no changes made", id), op)
	}

	var post models.Post
	err := r.storage.Sdb.Where("id = ?", id).First(&post).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &post, nil
}
