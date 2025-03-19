package services

import (
	"context"
	"fmt"
	"github.com/k1e1n04/video-streaming-sample/api/application/parameter"
	"github.com/k1e1n04/video-streaming-sample/api/domain/entities"
	"github.com/k1e1n04/video-streaming-sample/api/domain/repositories"
	"github.com/k1e1n04/video-streaming-sample/api/errors"
)

type VideoService struct {
	videoMetadataRepository repositories.VideoMetadataRepository
	videoStorageRepository  repositories.VideoStorageRepository
}

// NewVideoService is a constructor
func NewVideoService(
	videoMetadataRepository repositories.VideoMetadataRepository,
	videoStorageRepository repositories.VideoStorageRepository,
) VideoService {
	return VideoService{
		videoMetadataRepository: videoMetadataRepository,
		videoStorageRepository:  videoStorageRepository,
	}
}

// Register is a method to register video
func (v *VideoService) Register(ctx context.Context, p parameter.RegisterVideoParameter) (*string, error) {
	metadata, err := entities.NewVideoMetadataEntity(
		p.Title,
	)
	if err != nil {
		return nil, err
	}
	id := metadata.ID()
	if err := v.videoStorageRepository.Store(ctx, *id, p.Video); err != nil {
		return nil, err
	}
	if err := v.videoMetadataRepository.Register(ctx, *metadata); err != nil {
		return nil, err
	}
	idV := id.Value()
	return &idV, nil
}

// GetPresignedURLByVideoID is a method to get presigned URL by video ID
func (v *VideoService) GetPresignedURLByVideoID(ctx context.Context, p parameter.GetPresignedURLParameter) (*string, error) {
	id := entities.RestoreVideoID(p.VideoID)
	metadata, err := v.videoMetadataRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if metadata == nil {
		return nil, errors.NewNotFoundError(
			fmt.Sprintf("video metadata not found: %s", id.Value()),
			"video not found",
			nil,
		)
	}
	url, err := v.videoStorageRepository.GetPresignedURLByVideoID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &url, nil
}
