package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// ApplicationSetting is application setting
type ApplicationSetting struct {
	env                 string
	awsRegion           string
	videoBucketName     string
	thumbnailBucketName string
}

const LocalEnv = "local"

// NewApplicationSetting is generate ApplicationSetting
func NewApplicationSetting() *ApplicationSetting {
	if err := godotenv.Load(); err != nil {
		log.Printf(".env ファイルが存在しませんでした。")
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = LocalEnv
	}
	awsRegion := os.Getenv("AWS_REGION")
	if awsRegion == "" {
		awsRegion = "ap-northeast-1"
	}
	videoBucketName := os.Getenv("VIDEO_BUCKET_NAME")
	if videoBucketName == "" {
		videoBucketName = "local-video-bucket"
	}
	thumbnailBucketName := os.Getenv("THUMBNAIL_BUCKET_NAME")
	if thumbnailBucketName == "" {
		videoBucketName = "local-thumbnail-bucket"
	}

	return &ApplicationSetting{
		env:             env,
		awsRegion:       awsRegion,
		videoBucketName: videoBucketName,
	}
}

// Env is get env
func (a *ApplicationSetting) Env() string {
	return a.env
}

// AWSRegion is get aws region
func (a *ApplicationSetting) AWSRegion() string {
	return a.awsRegion
}

// VideoBucketName is get video bucket name
func (a *ApplicationSetting) VideoBucketName() string {
	return a.videoBucketName
}

// ThumbnailBucketName is get thumbnail bucket name
func (a *ApplicationSetting) ThumbnailBucketName() string {
	return a.thumbnailBucketName
}
