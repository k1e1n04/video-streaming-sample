package services

import (
	"context"
	"fmt"
	"github.com/k1e1n04/video-streaming-sample/api/shared/models"

	"github.com/k1e1n04/video-streaming-sample/api/video/application/dto"
	parameter2 "github.com/k1e1n04/video-streaming-sample/api/video/application/parameter"
	entities2 "github.com/k1e1n04/video-streaming-sample/api/video/domain/entities"
	repositories2 "github.com/k1e1n04/video-streaming-sample/api/video/domain/repositories"

	"github.com/k1e1n04/video-streaming-sample/api/utils"

	"github.com/k1e1n04/video-streaming-sample/api/errors"
)

type VideoService struct {
	videoMetadataRepository    repositories2.VideoMetadataRepository
	videoStorageRepository     repositories2.VideoStorageRepository
	thumbnailStorageRepository repositories2.ThumbnailStorageRepository
}

// NewVideoService is a constructor
func NewVideoService(
	videoMetadataRepository repositories2.VideoMetadataRepository,
	videoStorageRepository repositories2.VideoStorageRepository,
	thumbnailStorageRepository repositories2.ThumbnailStorageRepository,
) VideoService {
	return VideoService{
		videoMetadataRepository:    videoMetadataRepository,
		videoStorageRepository:     videoStorageRepository,
		thumbnailStorageRepository: thumbnailStorageRepository,
	}
}

// Register is a method to register video
func (v *VideoService) Register(ctx context.Context, p parameter2.RegisterVideoParameter) (*string, error) {
	userID := models.RestoreUserID(p.UserID)
	metadata, err := entities2.NewVideoMetadataEntity(
		p.Extension,
		p.Title,
		p.Description,
		p.ThumbnailExtension,
		p.Duration,
		p.Status,
		userID,
	)
	if err != nil {
		return nil, err
	}
	id := metadata.ID()
	if err := v.videoStorageRepository.Store(ctx, *id, p.Video, p.ThumbnailExtension); err != nil {
		return nil, err
	}
	if err := v.thumbnailStorageRepository.Store(ctx, *id, p.Thumbnail, p.ThumbnailExtension); err != nil {
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
	videoExtension := metadata.VideoExtension()
	url, err := v.videoStorageRepository.GetPresignedURLByVideoID(ctx, id, videoExtension.Value())
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
		thUrl, err := v.thumbnailStorageRepository.GetPresignedURLByVideoID(ctx, *(metadata.ID()), metadata.ThumbnailExtension().Value())
		if err != nil {
			return nil, err
		}
		dtos = append(dtos, dto.GetVideoPageDTO{
			ID:           metadata.ID().Value(),
			Title:        metadata.Title().Value(),
			Description:  metadata.Description().Value(),
			ThumbnailURL: thUrl,
			Views:        metadata.Views(),
			Likes:        metadata.Likes(),
			Duration:     metadata.Duration(),
			Status:       metadata.Status().String(),
			CreatedAt:    metadata.CreatedAt(),
		})
	}

	return utils.NewPageable(
		dtos,
		pageable.LastEvaluatedKey(),
	), nil
}
