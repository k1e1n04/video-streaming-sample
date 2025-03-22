package services

import (
	"context"
	"fmt"

	"github.com/k1e1n04/video-streaming-sample/api/video/application/dto"
	parameter2 "github.com/k1e1n04/video-streaming-sample/api/video/application/parameter"
	entities2 "github.com/k1e1n04/video-streaming-sample/api/video/domain/entities"
	repositories2 "github.com/k1e1n04/video-streaming-sample/api/video/domain/repositories"

	"github.com/k1e1n04/video-streaming-sample/api/utils"

	"github.com/k1e1n04/video-streaming-sample/api/errors"
)

type VideoService struct {
	videoMetadataRepository repositories2.VideoMetadataRepository
	videoStorageRepository  repositories2.VideoStorageRepository
}

// NewVideoService is a constructor
func NewVideoService(
	videoMetadataRepository repositories2.VideoMetadataRepository,
	videoStorageRepository repositories2.VideoStorageRepository,
) VideoService {
	return VideoService{
		videoMetadataRepository: videoMetadataRepository,
		videoStorageRepository:  videoStorageRepository,
	}
}

// Register is a method to register video
func (v *VideoService) Register(ctx context.Context, p parameter2.RegisterVideoParameter) (*string, error) {
	metadata, err := entities2.NewVideoMetadataEntity(
		p.Title,
	)
	if err != nil {
		return nil, err
	}
	id := metadata.ID()
	if err := v.videoStorageRepository.Store(ctx, *id, p.Video, "mp4"); err != nil {
		return nil, err
	}
	if err := v.videoMetadataRepository.Register(ctx, *metadata); err != nil {
		return nil, err
	}
	idV := id.Value()
	return &idV, nil
}

// GetPresignedURLByVideoID is a method to get presigned URL by video ID
func (v *VideoService) GetPresignedURLByVideoID(ctx context.Context, p parameter2.GetPresignedURLParameter) (*string, error) {
	id := entities2.RestoreVideoID(p.VideoID)
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

// GetVideoPage is a method to find page
func (v *VideoService) GetVideoPage(ctx context.Context, p parameter2.GetVideoPageParameter) (*utils.Pageable[dto.GetVideoPageDTO], error) {
	pageable, err := v.videoMetadataRepository.FindPage(ctx, p.Limit, p.LastEvaluatedKey)
	if err != nil {
		return nil, err
	}

	var dtos []dto.GetVideoPageDTO
	for _, metadata := range pageable.Content() {
		dtos = append(dtos, dto.GetVideoPageDTO{
			ID:        metadata.ID().Value(),
			Title:     metadata.Title().Value(),
			CreatedAt: metadata.CreatedAt(),
		})
	}

	return utils.NewPageable(
		dtos,
		pageable.LastEvaluatedKey(),
	), nil
}
