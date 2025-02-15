package repositories

import (
	"fmt"

	"github.com/gene-qxsi/Flexive/internal/models/orm_models"
	"github.com/gene-qxsi/Flexive/internal/storage"
)

type CommentRepo struct {
	storage *storage.Storage
}

func NewCommentRepo(storage *storage.Storage) *CommentRepo {
	return &CommentRepo{
		storage: storage,
	}
}

func (r *CommentRepo) CreateComment(comment *orm_models.Comment) (*orm_models.Comment, error) {
	const op = "internal/api/repositories/comment_repo.go/CreateComment()"

	err := r.storage.Sdb.Create(comment).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return comment, nil
}

func (r *CommentRepo) GetComment(id int) (*orm_models.Comment, error) {
	const op = "internal/api/repositories/comment_repo.go/GetComment()"

	var comment orm_models.Comment
	err := r.storage.Sdb.First(&comment, id).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &comment, nil
}

func (r *CommentRepo) GetComments() ([]orm_models.Comment, error) {
	const op = "internal/api/repositories/comment_repo.go/GetComments()"

	var comments []orm_models.Comment
	err := r.storage.Sdb.Find(&comments).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return comments, nil
}

func (r *CommentRepo) DeleteComment(id int) error {
	const op = "internal/api/repositories/comment_repo.go/DeleteComment()"

	result := r.storage.Sdb.Delete(&orm_models.Comment{}, id)
	if result.Error != nil {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-2: %s. ПУТЬ: %s", fmt.Sprintf("comment with ID %d not found", id), op)
	}

	return nil
}

func (r *CommentRepo) UpdateComment(id int, values map[string]interface{}) (*orm_models.Comment, error) {
	const op = "internal/api/repositories/comment_repo.go/UpdateComment()"

	result := r.storage.Sdb.Model(&orm_models.Comment{}).Where("id = ?", id).Updates(values)
	if result.Error != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", fmt.Sprintf("comment with ID %d not found or no changes made", id), op)
	}

	var comment orm_models.Comment
	err := r.storage.Sdb.Where("id = ?", id).First(&comment).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &comment, nil
}
