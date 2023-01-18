package repository

import (
	"gorm.io/gorm"

	"golang-rate-limitter/internal/blog/model"
)

type BlogRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{db: db}
}

func (r *BlogRepository) CreatePost(blog model.BlogModel) error {
	if err := r.db.Create(&blog).Error; err != nil {
		return err
	}
	return nil
}
