package models

type URLMapping struct {
	ID       uint   `gorm:"primaryKey"`
	ShortURL string `gorm:"uniqueIndex"`
	LongURL  string
}
