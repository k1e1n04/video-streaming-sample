package dto

import "time"

// GetVideoPageDTO is a DTO for GetVideoPage
type GetVideoPageDTO struct {
	// ID is a video ID
	ID string
	// Title is a title
	Title string
	// CreatedAt is a created at
	CreatedAt time.Time
}
