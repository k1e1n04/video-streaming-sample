package dto

import "time"

// GetVideoPageDTO is a DTO for GetVideoPage
type GetVideoPageDTO struct {
	// ID is a video ID
	ID string
	// Title is a title
	Title string
	// Description is a description
	Description string
	// ThumbnailURL is a thumbnail URL
	ThumbnailURL string
	// Views is a views
	Views int64
	// Likes is a likes
	Likes int64
	// Duration is a duration
	Duration int64
	// Status is a status
	Status string
	// CreatedAt is a created at
	CreatedAt time.Time
}
