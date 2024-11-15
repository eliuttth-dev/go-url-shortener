package main

import (
	"time"
)

type URL struct {
	ID 		uint		`json:"id" gorm:"primaryKey"`
	OriginalURL	string		`json:"original_url"`
	ShortURL	string		`json:"short_url"`
	ClickCoun	int		`json:"click_count"`
	ExpiresAt	time.Time	`json:"expires_at"`
}
