package repositories

import (
	"fmt"

	"github.com/gene-qxsi/Flexive/internal/models/orm_models"
	"github.com/gene-qxsi/Flexive/internal/storage"
)

type ReactionRepo struct {
	storage *storage.Storage
}

func NewReactionRepo(storage *storage.Storage) *ReactionRepo {
	return &ReactionRepo{
		storage: storage,
	}
}

func (r *ReactionRepo) CreateReaction(reaaction *orm_models.Reaction) (*orm_models.Reaction, error) {
	const op = "internal/api/repositories/reaction_repo.go/CreateReaction()"

	err := r.storage.Sdb.Create(reaaction).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return reaaction, nil
}

func (r *ReactionRepo) GetReaction(userID, postID int) (*orm_models.Reaction, error) {
	const op = "internal/api/repositories/reaction_repo.go/GetReaction()"

	var reaction orm_models.Reaction
	err := r.storage.Sdb.First(&reaction, userID, postID).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &reaction, nil
}

func (r *ReactionRepo) GetReactions() ([]orm_models.Reaction, error) {
	const op = "internal/api/repositories/reaction_repo.go/GetReactions()"

	var reactions []orm_models.Reaction
	err := r.storage.Sdb.Find(&reactions).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return reactions, nil
}

func (r *ReactionRepo) DeleteReaction(userID, postID int) error {
	const op = "internal/api/repositories/reaction_repo.go/DeleteReaction()"

	result := r.storage.Sdb.Delete(&orm_models.Reaction{}, userID, postID)
	if result.Error != nil {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-2: %s. ПУТЬ: %s",
			fmt.Sprintf("reaction with UserID %d AND PostID %d not found", userID, postID), op)
	}

	return nil
}

func (r *ReactionRepo) UpdateReaction(userID, postID int, values map[string]interface{}) (*orm_models.Reaction, error) {
	const op = "internal/api/repositories/reaction_repo.go/UpdateReaction()"

	result := r.storage.Sdb.Model(&orm_models.Reaction{}).Where("user_id = ? AND post_id = ?", userID, postID).Updates(values)
	if result.Error != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s",
			fmt.Sprintf("reaction with UserID %d AND PostID %d not found or no changes made", userID, postID), op)
	}
	var reaction *orm_models.Reaction
	err := r.storage.Sdb.Where("user_id = ?", userID).Where("post_id = ?", postID).First(&reaction).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op)
	}

	return reaction, nil
}
