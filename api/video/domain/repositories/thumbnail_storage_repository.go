package repositories

import (
	"bytes"
	"context"
	"github.com/k1e1n04/video-streaming-sample/api/video/domain/entities"
)

type ThumbnailStorageRepository interface {
	// Store is a method to store thumbnail
	Store(ctx context.Context, videoID entities.VideoID, thumbnail *bytes.Reader, extension string) error
	// GetPresignedURLByVideoID is a method to get unsigned URL by video ID
	GetPresignedURLByVideoID(ctx context.Context, videoID entities.VideoID, extension string) (string, error)
}
