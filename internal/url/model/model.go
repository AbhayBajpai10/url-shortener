package model

import ("time");

type URL struct {
	ID          int
	UserID      int
	OriginalURL string
	ShortCode   string
	Clicks      int
	CreatedAt   time.Time
}