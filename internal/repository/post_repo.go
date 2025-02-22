package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gene-qxsi/Flexive/configs"
	"github.com/gene-qxsi/Flexive/internal/repository/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	rankPosts = "posts:ranking"
)

type PostRepo struct {
	db             *gorm.DB
	client         *redis.Client
	rankingPostTTL time.Duration
}

func NewPostRepo(db *gorm.DB, client *redis.Client, conf *configs.Config) *PostRepo {
	return &PostRepo{
		db:             db,
		client:         client,
		rankingPostTTL: conf.PostsRankingTTL,
	}
}

func (r *PostRepo) CreatePost(post *models.Post) (*models.Post, error) {
	const op = "internal/api/repositories/post_repo.go/CreatePost()"

	err := r.db.Create(post).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return post, nil
}

func (r *PostRepo) GetPost(id int) (*models.Post, error) {
	const op = "internal/api/repositories/post_repo.go/GetPost()"

	var post models.Post
	err := r.db.First(&post, id).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &post, nil
}

func (r *PostRepo) GetPosts() ([]models.Post, error) {
	const op = "internal/api/repositories/post_repo.go/GetPosts()"

	var posts []models.Post
	err := r.db.Find(&posts).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return posts, nil
}

func (r *PostRepo) DeletePost(id int) error {
	const op = "internal/api/repositories/post_repo.go/DeletePost()"

	result := r.db.Delete(&models.Post{}, id)
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

	result := r.db.Model(&models.Post{}).Where("id = ?", id).Updates(values)
	if result.Error != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", fmt.Sprintf("post with ID %d not found or no changes made", id), op)
	}

	var post models.Post
	err := r.db.Where("id = ?", id).First(&post).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &post, nil
}

func (r *PostRepo) IncrPost(ctx context.Context, member string) error {
	const op = "internal/cache/redis_db.go/CachePost()"

	_, err := r.client.ZIncrBy(ctx, rankPosts, 1, member).Result()
	if err != nil {
		return fmt.Errorf("❌ КЕШ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	_, err = r.client.Expire(ctx, rankPosts, r.rankingPostTTL).Result()
	if err != nil {
		return fmt.Errorf("❌ КЕШ-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op)
	}

	return nil
}

func (r *PostRepo) GetTopPostsIDs(ctx context.Context, limit int) ([]int, error) {
	const op = "internal/cache/redis_db.go/GetTopN()"

	idS, err := r.client.ZRevRange(ctx, rankPosts, 0, int64(limit-1)).Result()
	if err != nil {
		return nil, fmt.Errorf("❌ КЕШ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	var idIs []int
	for _, id := range idS {
		idI, err := strconv.Atoi(id)
		if err != nil {
			return nil, fmt.Errorf("❌ КЕШ-ОШИБКА-2: %s. ПУТЬ: %s", err, op)
		}
		idIs = append(idIs, idI)
	}

	return idIs, nil
}
