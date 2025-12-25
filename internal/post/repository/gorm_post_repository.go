package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
	"github.com/mxilia/Quonet-backend/internal/transaction"
	"gorm.io/gorm"
)

type GormPostRepository struct {
	db *gorm.DB
}

func NewGormPostRepository(db *gorm.DB) PostRepository {
	return &GormPostRepository{db: db}
}

func (r *GormPostRepository) Save(ctx context.Context, post *entities.Post) error {
	tx := transaction.GetTx(ctx, r.db)
	if err := tx.Create(post).Error; err != nil {
		return err
	}
	return nil
}

/* No private posts involved */
func (r *GormPostRepository) Find(authorID uuid.UUID, threadID uuid.UUID, title string, offset int, limit int) ([]*entities.Post, error) {
	query := r.db.Preload("Author").Preload("Thread").Where("is_private = ?", false)
	if authorID != uuid.Nil {
		query = query.Where("author_id = ?", authorID)
	}

	if threadID != uuid.Nil {
		query = query.Where("thread_id = ?", threadID)
	}

	if title != "" {
		query = query.Where("title = ?", title)
	}

	var postsValue []entities.Post
	if err := query.Limit(limit).Offset(offset).Find(&postsValue).Error; err != nil {
		return nil, err
	}

	posts := make([]*entities.Post, len(postsValue))
	for i := range postsValue {
		posts[i] = &postsValue[i]
	}
	return posts, nil
}

func (r *GormPostRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.Post, error) {
	tx := transaction.GetTx(ctx, r.db)

	var post entities.Post
	if err := tx.Preload("Author").Preload("Thread").Where("is_private = ?", false).First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *GormPostRepository) FindTopLiked(authorID uuid.UUID, threadID uuid.UUID, title string, limit int) ([]*entities.Post, error) {
	query := r.db.Preload("Author").Preload("Thread").Where("is_private = ?", false)
	if authorID != uuid.Nil {
		query = query.Where("author_id = ?", authorID)
	}

	if threadID != uuid.Nil {
		query = query.Where("thread_id = ?", threadID)
	}

	if title != "" {
		query = query.Where("title = ?", title)
	}

	var postsValue []entities.Post
	if err := query.Limit(limit).Order("like_count DESC").Find(&postsValue).Error; err != nil {
		return nil, err
	}

	posts := make([]*entities.Post, len(postsValue))
	for i := range postsValue {
		posts[i] = &postsValue[i]
	}
	return posts, nil
}

/* Private posts involved */
func (r *GormPostRepository) FindPrivate(authorID uuid.UUID, threadID uuid.UUID, title string, offset int, limit int) ([]*entities.Post, error) {
	query := r.db.Preload("Author").Preload("Thread").Where("is_private = ?", true)
	if authorID != uuid.Nil {
		query = query.Where("author_id = ?", authorID)
	}

	if threadID != uuid.Nil {
		query = query.Where("thread_id = ?", threadID)
	}

	if title != "" {
		query = query.Where("title = ?", title)
	}

	var postsValue []entities.Post
	if err := query.Limit(limit).Offset(offset).Find(&postsValue).Error; err != nil {
		return nil, err
	}

	posts := make([]*entities.Post, len(postsValue))
	for i := range postsValue {
		posts[i] = &postsValue[i]
	}
	return posts, nil
}

func (r *GormPostRepository) FindPrivateByID(id uuid.UUID) (*entities.Post, error) {
	var post entities.Post
	if err := r.db.Preload("Author").Preload("Thread").Where("is_private = ?", true).First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *GormPostRepository) Count(isPrivate bool, authorID uuid.UUID, threadID uuid.UUID, title string) (int64, error) {
	query := r.db.Model(&entities.Post{}).Where("is_private = ?", isPrivate)
	if authorID != uuid.Nil {
		query = query.Where("author_id = ?", authorID)
	}

	if threadID != uuid.Nil {
		query = query.Where("thread_id = ?", threadID)
	}

	if title != "" {
		query = query.Where("title = ?", title)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

func (r *GormPostRepository) FindNoFilterByID(id uuid.UUID) (*entities.Post, error) {
	var post entities.Post
	if err := r.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *GormPostRepository) Patch(ctx context.Context, id uuid.UUID, post *entities.Post) error {
	tx := transaction.GetTx(ctx, r.db)

	result := tx.Model(&entities.Post{}).Where("id = ?", id).Updates(post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormPostRepository) Delete(ctx context.Context, id uuid.UUID) error {
	tx := transaction.GetTx(ctx, r.db)

	result := tx.Delete(&entities.Post{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
