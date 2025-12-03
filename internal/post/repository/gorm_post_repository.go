package repository

import (
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
	"gorm.io/gorm"
)

type GormPostRepository struct {
	db *gorm.DB
}

func NewGormPostRepository(db *gorm.DB) PostRepository {
	return &GormPostRepository{db: db}
}

func (r *GormPostRepository) Save(post *entities.Post) error {
	if err := r.db.Create(post).Error; err != nil {
		return err
	}
	return nil
}

/* No private posts involved */
func (r *GormPostRepository) FindAll() ([]*entities.Post, error) {
	var postsValue []entities.Post
	if err := r.db.Preload("Author").Where("is_private = ?", false).Find(&postsValue).Error; err != nil {
		return nil, err
	}

	posts := make([]*entities.Post, len(postsValue))
	for i := range postsValue {
		posts[i] = &postsValue[i]
	}
	return posts, nil
}

func (r *GormPostRepository) FindByAuthorID(id uuid.UUID) ([]*entities.Post, error) {
	var postsValue []entities.Post
	if err := r.db.Preload("Author").Where("is_private = ?", false).Where("author_id = ?", id).Find(&postsValue).Error; err != nil {
		return nil, err
	}

	posts := make([]*entities.Post, len(postsValue))
	for i := range postsValue {
		posts[i] = &postsValue[i]
	}
	return posts, nil
}

func (r *GormPostRepository) FindByThreadID(id uuid.UUID) ([]*entities.Post, error) {
	var postsValue []entities.Post
	if err := r.db.Preload("Author").Where("is_private = ?", false).Where("thread_id = ?", id).Find(&postsValue).Error; err != nil {
		return nil, err
	}

	posts := make([]*entities.Post, len(postsValue))
	for i := range postsValue {
		posts[i] = &postsValue[i]
	}
	return posts, nil
}

func (r *GormPostRepository) FindByID(id uuid.UUID) (*entities.Post, error) {
	var post entities.Post
	if err := r.db.Preload("Author").Where("is_private = ?", false).First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *GormPostRepository) FindByTitle(title string) (*entities.Post, error) {
	var post entities.Post
	if err := r.db.Preload("Author").Where("is_private = ?", false).Where("title = ?", title).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

/* Private posts involved */
func (r *GormPostRepository) FindAllCoverPrivate() ([]*entities.Post, error) {
	var postsValue []entities.Post
	if err := r.db.Find(&postsValue).Error; err != nil {
		return nil, err
	}

	posts := make([]*entities.Post, len(postsValue))
	for i := range postsValue {
		posts[i] = &postsValue[i]
	}
	return posts, nil
}

func (r *GormPostRepository) FindAllPrivate() ([]*entities.Post, error) {
	var postsValue []entities.Post
	if err := r.db.Where("is_private = ?", true).Find(&postsValue).Error; err != nil {
		return nil, err
	}

	posts := make([]*entities.Post, len(postsValue))
	for i := range postsValue {
		posts[i] = &postsValue[i]
	}
	return posts, nil
}

func (r *GormPostRepository) FindCoverPrivateByAuthorID(id uuid.UUID) ([]*entities.Post, error) {
	var postsValue []entities.Post
	if err := r.db.Where("author_id = ?", id).Find(&postsValue).Error; err != nil {
		return nil, err
	}

	posts := make([]*entities.Post, len(postsValue))
	for i := range postsValue {
		posts[i] = &postsValue[i]
	}
	return posts, nil
}

func (r *GormPostRepository) FindPrivateByAuthorID(id uuid.UUID) ([]*entities.Post, error) {
	var postsValue []entities.Post
	if err := r.db.Where("is_private = ?", true).Where("author_id = ?", id).Find(&postsValue).Error; err != nil {
		return nil, err
	}

	posts := make([]*entities.Post, len(postsValue))
	for i := range postsValue {
		posts[i] = &postsValue[i]
	}
	return posts, nil
}

func (r *GormPostRepository) FindCoverPrivateByThreadID(id uuid.UUID) ([]*entities.Post, error) {
	var postsValue []entities.Post
	if err := r.db.Where("thread_id = ?", id).Find(&postsValue).Error; err != nil {
		return nil, err
	}

	posts := make([]*entities.Post, len(postsValue))
	for i := range postsValue {
		posts[i] = &postsValue[i]
	}
	return posts, nil
}

func (r *GormPostRepository) FindPrivateByThreadID(id uuid.UUID) ([]*entities.Post, error) {
	var postsValue []entities.Post
	if err := r.db.Where("is_private = ?", true).Where("thread_id = ?", id).Find(&postsValue).Error; err != nil {
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
	if err := r.db.Where("is_private = ?", true).First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *GormPostRepository) FindPrivateByTitle(title string) (*entities.Post, error) {
	var post entities.Post
	if err := r.db.Where("is_private = ?", true).Where("title = ?", title).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *GormPostRepository) Patch(id uuid.UUID, post *entities.Post) error {
	result := r.db.Model(&entities.Post{}).Where("id = ?", id).Updates(post)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *GormPostRepository) Delete(id uuid.UUID) error {
	result := r.db.Delete(&entities.Post{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
