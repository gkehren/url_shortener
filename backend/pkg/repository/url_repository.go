package repository

import (
	"url_shortener/pkg/models"

	"gorm.io/gorm"
)

type URLRepository struct {
	DB *gorm.DB
}

func NewURLRepository(db *gorm.DB) *URLRepository {
	return &URLRepository{DB: db}
}

func (r *URLRepository) CreateURLMapping(shortURL, longURL string) error {
	mapping := models.URLMapping{ShortURL: shortURL, LongURL: longURL}
	return r.DB.Create(&mapping).Error
}

func (r *URLRepository) GetURLMappingByLongURL(longURL string) (models.URLMapping, error) {
	var mapping models.URLMapping
	err := r.DB.Where("long_url = ?", longURL).First(&mapping).Error
	return mapping, err
}

func (r *URLRepository) GetURLMappingByShortURL(shortURL string) (models.URLMapping, error) {
	var mapping models.URLMapping
	err := r.DB.Where("short_url = ?", shortURL).First(&mapping).Error
	return mapping, err
}
