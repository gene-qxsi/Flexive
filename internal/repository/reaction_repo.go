package repository

import (
	"fmt"

	"github.com/gene-qxsi/Flexive/internal/repository/models"
	"gorm.io/gorm"
)

type ReactionRepo struct {
	db *gorm.DB
}

func NewReactionRepo(db *gorm.DB) *ReactionRepo {
	return &ReactionRepo{
		db: db,
	}
}

func (r *ReactionRepo) CreateReaction(reaaction *models.Reaction) (*models.Reaction, error) {
	const op = "internal/api/repositories/reaction_repo.go/CreateReaction()"

	err := r.db.Create(reaaction).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return reaaction, nil
}

func (r *ReactionRepo) GetReaction(userID, postID int) (*models.Reaction, error) {
	const op = "internal/api/repositories/reaction_repo.go/GetReaction()"

	var reaction models.Reaction
	err := r.db.First(&reaction, userID, postID).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &reaction, nil
}

func (r *ReactionRepo) GetReactions() ([]models.Reaction, error) {
	const op = "internal/api/repositories/reaction_repo.go/GetReactions()"

	var reactions []models.Reaction
	err := r.db.Find(&reactions).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return reactions, nil
}

func (r *ReactionRepo) DeleteReaction(userID, postID int) error {
	const op = "internal/api/repositories/reaction_repo.go/DeleteReaction()"

	result := r.db.Delete(&models.Reaction{}, userID, postID)
	if result.Error != nil {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-2: %s. ПУТЬ: %s",
			fmt.Sprintf("reaction with UserID %d AND PostID %d not found", userID, postID), op)
	}

	return nil
}

func (r *ReactionRepo) UpdateReaction(userID, postID int, values map[string]interface{}) (*models.Reaction, error) {
	const op = "internal/api/repositories/reaction_repo.go/UpdateReaction()"

	result := r.db.Model(&models.Reaction{}).Where("user_id = ? AND post_id = ?", userID, postID).Updates(values)
	if result.Error != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s",
			fmt.Sprintf("reaction with UserID %d AND PostID %d not found or no changes made", userID, postID), op)
	}

	var reaction *models.Reaction
	err := r.db.Where("user_id = ?", userID).Where("post_id = ?", postID).First(&reaction).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op)
	}

	return reaction, nil
}
