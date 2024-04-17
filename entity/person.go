package entity

import "time"

// Person - store data related to users
type Person struct {
	ID               int64     `json:"id"`
	TelegramID       int64     `json:"telegram_id"`
	FirstName        string    `josn:"first_name"`
	LastName         string    `json:"last_name"`
	Language         string    `json:"language"`
	TelegramLanguage string    `json:"telegram_language"`
	MainPasswordHash string    `json:"main_pass_hash"`
	FakePasswordHash string    `json:"fake_pass_hash"`
	LocationLat      float64   `json:"location_lat"`
	LocationLon      float64   `json:"location_lon"`
	IsAdmin          bool      `json:"is_admin"`
	CreatedAt        time.Time `json:"created_at"`
}
