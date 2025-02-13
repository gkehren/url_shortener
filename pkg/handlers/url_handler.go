package handlers

import (
	"url_shortener/pkg/repository"
	"url_shortener/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type URLHandler struct {
	repo *repository.URLRepository
}

func NewURLHandler(db *gorm.DB) *URLHandler {
	return &URLHandler{repo: repository.NewURLRepository(db)}
}

func (h *URLHandler) ShortenURL(c *gin.Context) {
	var input struct {
		LongURL string `json:"long_url"`
	}
	if c.BindJSON(&input) == nil {
		existingMapping, err := h.repo.GetURLMappingByLongURL(input.LongURL)
		if err == nil {
			c.JSON(200, gin.H{"short_url": existingMapping.ShortURL})
			return
		}

		shortURL := utils.GenerateShortURL()
		err = h.repo.CreateURLMapping(shortURL, input.LongURL)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create URL mapping"})
			return
		}
		c.JSON(200, gin.H{"short_url": shortURL})
	}
}

func (h *URLHandler) ResolveURL(c *gin.Context) {
	shortURL := c.Param("shortURL")
	mapping, err := h.repo.GetURLMappingByShortURL(shortURL)
	if err != nil {
		c.String(404, "URL not found")
		return
	}
	c.Redirect(302, mapping.LongURL)
}
