package parameter

import (
	"bytes"
	"github.com/k1e1n04/video-streaming-sample/api/video/domain/entities"
)

// RegisterVideoParameter is a parameter for RegisterVideo
type RegisterVideoParameter struct {
	// UserID is a user ID
	UserID string
	// Title is a title
	Title string
	// Description is a description
	Description string
	// Extension is an extension of the video
	Extension string
	// Duration is a duration of the video
	Duration int64
	// ThumbnailExtension is an extension of the thumbnail
	ThumbnailExtension string
	// Thumbnail is a thumbnail
	Thumbnail *bytes.Reader
	// Status is a status of the video
	Status entities.VideoStatus
	// Video is a video
	Video *bytes.Reader
}
