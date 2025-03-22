package repositories

import (
	"context"

	entities2 "github.com/k1e1n04/video-streaming-sample/api/domain/entities"
)

// VideoMetadataRepository is a video metadata repository
type VideoMetadataRepository interface {
	// Register is a method to register video metadata
	Register(ctx context.Context, video entities2.VideoMetadataEntity) error
	// FindByID is a method to find video metadata by id
	FindByID(ctx context.Context, id entities2.VideoID) (*entities2.VideoMetadataEntity, error)
}
