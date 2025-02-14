package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"url_shortener/pkg/models"
	"url_shortener/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.URLMapping{})
	return db
}

func TestShortenURL(t *testing.T) {
	db := setupTestDB()
	handler := NewURLHandler(db)
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/shorten", handler.ShortenURL)

	longURL := "https://example.com"
	body, _ := json.Marshal(gin.H{"long_url": longURL})
	req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.NotEmpty(t, response["short_url"])

	// Verify the URL mapping is stored in the database
	var mapping models.URLMapping
	db.Where("long_url = ?", longURL).First(&mapping)
	assert.Equal(t, longURL, mapping.LongURL)
	assert.Equal(t, response["short_url"], mapping.ShortURL)
}

func TestResolveURL(t *testing.T) {
	db := setupTestDB()
	handler := NewURLHandler(db)
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/:shortURL", handler.ResolveURL)

	shortURL := utils.GenerateShortURL()
	longURL := "https://example.com"
	db.Create(&models.URLMapping{ShortURL: shortURL, LongURL: longURL})

	req, _ := http.NewRequest("GET", "/"+shortURL, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, longURL, w.Header().Get("Location"))
}
