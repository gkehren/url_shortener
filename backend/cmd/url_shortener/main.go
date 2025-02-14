package main

import (
	"url_shortener/internal/config"
	"url_shortener/pkg/handlers"
	repository "url_shortener/pkg/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("urls.db"), &gorm.Config{})
	if err != nil {
		panic("faile to connect database")
	}
	db.AutoMigrate(&repository.URLMapping{})
}

func main() {
	config.LoadConfig()
	initDB()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	handler := handlers.NewURLHandler(db)

	r.POST("/shorten", handler.ShortenURL)
	r.GET("/:shortURL", handler.ResolveURL)
	r.Run(":8080")
}
