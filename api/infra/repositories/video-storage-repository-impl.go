package repositories

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	entities2 "github.com/k1e1n04/video-streaming-sample/api/domain/entities"
	"github.com/k1e1n04/video-streaming-sample/api/domain/repositories"
	"github.com/k1e1n04/video-streaming-sample/api/env"
)

// VideoStorageRepositoryImpl is a struct to store video storage repository implementation
type VideoStorageRepositoryImpl struct {
	// s3Client is a s3 client
	s3Client *s3.Client
	// s3Uploader is a s3 uploader
	s3Uploader *manager.Uploader
	// setting is a application setting
	setting *env.ApplicationSetting
}

// NewVideoStorageRepositoryImpl is a function to create a new VideoStorageRepositoryImpl
func NewVideoStorageRepositoryImpl(
	s3Client *s3.Client,
	s3Uploader *manager.Uploader,
	setting *env.ApplicationSetting,
) repositories.VideoStorageRepository {
	return &VideoStorageRepositoryImpl{
		s3Client:   s3Client,
		s3Uploader: s3Uploader,
		setting:    setting,
	}
}

// Store is a method to store video
func (v *VideoStorageRepositoryImpl) Store(ctx context.Context, videoID entities2.VideoID, video *bytes.Reader, extension string) error {
	_, err := v.s3Uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(v.setting.VideoBucketName()),
		ContentType: aws.String("video/mp4"),
		Key:         aws.String(videoID.Value() + "." + extension),
		Body:        video,
	})
	if err != nil {
		return err
	}
	return nil
}

// GetPresignedURLByVideoID is a method to get presigned URL by video ID
func (v *VideoStorageRepositoryImpl) GetPresignedURLByVideoID(ctx context.Context, videoID entities2.VideoID) (string, error) {
	presignedURL, err := s3.NewPresignClient(v.s3Client).PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(v.setting.VideoBucketName()),
		Key:    aws.String(videoID.Value() + ".mp4"),
	})
	if err != nil {
		return "", err
	}
	return presignedURL.URL, nil
}
