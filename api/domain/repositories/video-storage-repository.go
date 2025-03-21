package repositories

import (
	"bytes"
	"context"
	"github.com/k1e1n04/video-streaming-sample/api/domain/entities"
)

// VideoStorageRepository is a video storage repository
type VideoStorageRepository interface {
	// Store is a method to store video
	Store(ctx context.Context, videoID entities.VideoID, video *bytes.Reader, extension string) error
	// GetPresignedURLByVideoID is a method to get unsigned URL by video ID
	GetPresignedURLByVideoID(ctx context.Context, videoID entities.VideoID) (string, error)
}
