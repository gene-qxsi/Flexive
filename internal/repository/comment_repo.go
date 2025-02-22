package repository

import (
	"fmt"

	"github.com/gene-qxsi/Flexive/internal/repository/models"
	"gorm.io/gorm"
)

type CommentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) *CommentRepo {
	return &CommentRepo{
		db: db,
	}
}

func (r *CommentRepo) CreateComment(comment *models.Comment) (*models.Comment, error) {
	const op = "internal/api/repositories/comment_repo.go/CreateComment()"

	err := r.db.Create(comment).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return comment, nil
}

func (r *CommentRepo) GetComment(id int) (*models.Comment, error) {
	const op = "internal/api/repositories/comment_repo.go/GetComment()"

	var comment models.Comment
	err := r.db.First(&comment, id).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &comment, nil
}

func (r *CommentRepo) GetComments() ([]models.Comment, error) {
	const op = "internal/api/repositories/comment_repo.go/GetComments()"

	var comments []models.Comment
	err := r.db.Find(&comments).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return comments, nil
}

func (r *CommentRepo) DeleteComment(id int) error {
	const op = "internal/api/repositories/comment_repo.go/DeleteComment()"

	result := r.db.Delete(&models.Comment{}, id)
	if result.Error != nil {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-2: %s. ПУТЬ: %s", fmt.Sprintf("comment with ID %d not found", id), op)
	}

	return nil
}

func (r *CommentRepo) UpdateComment(id int, values map[string]interface{}) (*models.Comment, error) {
	const op = "internal/api/repositories/comment_repo.go/UpdateComment()"

	result := r.db.Model(&models.Comment{}).Where("id = ?", id).Updates(values)
	if result.Error != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", fmt.Sprintf("comment with ID %d not found or no changes made", id), op)
	}

	var comment models.Comment
	err := r.db.Where("id = ?", id).First(&comment).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &comment, nil
}
