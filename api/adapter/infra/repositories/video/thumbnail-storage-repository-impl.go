package repositories

import (
	"bytes"
	"context"

	entities2 "github.com/k1e1n04/video-streaming-sample/api/video/domain/entities"
	"github.com/k1e1n04/video-streaming-sample/api/video/domain/repositories"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/k1e1n04/video-streaming-sample/api/env"
)

// ThumbnailStorageRepositoryImpl is a struct to store thumbnail storage repository implementation
type ThumbnailStorageRepositoryImpl struct {
	// s3Client is a s3 client
	s3Client *s3.Client
	// s3Uploader is a s3 uploader
	s3Uploader *manager.Uploader
	// setting is a application setting
	setting *env.ApplicationSetting
}

// NewThumbnailStorageRepositoryImpl is a function to create a new ThumbnailStorageRepositoryImpl
func NewThumbnailStorageRepositoryImpl(
	s3Client *s3.Client,
	s3Uploader *manager.Uploader,
	setting *env.ApplicationSetting,
) repositories.ThumbnailStorageRepository {
	return &ThumbnailStorageRepositoryImpl{
		s3Client:   s3Client,
		s3Uploader: s3Uploader,
		setting:    setting,
	}
}

// Store is a method to store video
func (v *ThumbnailStorageRepositoryImpl) Store(ctx context.Context, videoID entities2.VideoID, video *bytes.Reader, extension string) error {
	_, err := v.s3Uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(v.setting.ThumbnailBucketName()),
		Key:    aws.String(videoID.Value() + "." + extension),
		Body:   video,
	})
	if err != nil {
		return err
	}
	return nil
}

// GetPresignedURLByVideoID is a method to get presigned URL by video ID
func (v *ThumbnailStorageRepositoryImpl) GetPresignedURLByVideoID(ctx context.Context, videoID entities2.VideoID, extension string) (string, error) {
	presignedURL, err := s3.NewPresignClient(v.s3Client).PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(v.setting.ThumbnailBucketName()),
		Key:    aws.String(videoID.Value() + "." + extension),
	})
	if err != nil {
		return "", err
	}
	return presignedURL.URL, nil
}
