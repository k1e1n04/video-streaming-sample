package di

import (
	"context"
	"github.com/k1e1n04/video-streaming-sample/api/infra/repositories/video"
	"github.com/k1e1n04/video-streaming-sample/api/video/application/services"
	repositories2 "github.com/k1e1n04/video-streaming-sample/api/video/domain/repositories"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/k1e1n04/video-streaming-sample/api/adapter/controllers"
	"github.com/k1e1n04/video-streaming-sample/api/env"
	"go.uber.org/dig"
)

// Init is a function to initialize dependencies
func Init() *dig.Container {
	container := dig.New()

	setting := env.NewApplicationSetting()
	initSetting(container, setting)
	initClient(container, setting)
	initRepositories(container)
	initServices(container)
	initControllers(container)

	return container
}

// initSetting is a function to initialize settings
func initSetting(container *dig.Container, setting *env.ApplicationSetting) {
	err := container.Provide(func() *env.ApplicationSetting {
		return setting
	})
	if err != nil {
		panic(err)
	}
}

// initClient is a function to initialize clients
func initClient(container *dig.Container, setting *env.ApplicationSetting) {
	cfg := aws.Config{
		Region: setting.AWSRegion(),
	}
	if setting.Env() == env.LocalEnv {
		var err error
		cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
				os.Getenv("MINIO_ACCESS_KEY"), // access key
				os.Getenv("MINIO_SECRET_KEY"), // secret key
				"",
			)),
		)
		if err != nil {
			panic(err)
		}
	}

	err := container.Provide(func() *s3.Client {
		var client *s3.Client
		if setting.Env() == env.LocalEnv {
			client = s3.NewFromConfig(cfg, func(o *s3.Options) {
				o.BaseEndpoint = aws.String("http://localhost:9000")
				o.UsePathStyle = true
			})
		} else {
			client = s3.NewFromConfig(cfg)
		}
		return client
	})
	if err != nil {
		panic(err)
	}

	err = container.Provide(func() *dynamodb.Client {
		var client *dynamodb.Client
		if setting.Env() == env.LocalEnv {
			client = dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
				o.BaseEndpoint = aws.String("http://localhost:8000")
			})
		} else {
			client = dynamodb.NewFromConfig(cfg)
		}
		return client
	})
	if err != nil {
		panic(err)
	}

	err = container.Provide(func(s3Client *s3.Client) *manager.Uploader {
		return manager.NewUploader(s3Client, func(u *manager.Uploader) {
			u.PartSize = 5 * 1024 * 1024
			u.Concurrency = 5
		})
	})
	if err != nil {
		panic(err)
	}

	err = container.Provide(func(s3Client *s3.Client) *s3.PresignClient {
		return s3.NewPresignClient(s3Client)
	})
	if err != nil {
		panic(err)
	}
}

// initRepositories is a function to initialize repositories
func initRepositories(container *dig.Container) {
	err := container.Provide(func(dynamodb *dynamodb.Client) repositories2.VideoMetadataRepository {
		return video.NewVideoMetadataRepositoryImpl(dynamodb)
	})
	if err != nil {
		panic(err)
	}

	err = container.Provide(func(
		s3Client *s3.Client,
		s3Uploader *manager.Uploader,
		setting *env.ApplicationSetting,
	) repositories2.VideoStorageRepository {
		return video.NewVideoStorageRepositoryImpl(s3Client, s3Uploader, setting)
	})
	if err != nil {
		panic(err)
	}
}

// initServices is a function to initialize services
func initServices(container *dig.Container) {
	err := container.Provide(func(
		videoMetadataRepository repositories2.VideoMetadataRepository,
		videoStorageRepository repositories2.VideoStorageRepository,
	) services.VideoService {
		return services.NewVideoService(videoMetadataRepository, videoStorageRepository)
	})
	if err != nil {
		panic(err)
	}
}

// initControllers is a function to initialize controllers
func initControllers(container *dig.Container) {
	err := container.Provide(func(
		videoService services.VideoService,
	) controllers.VideoController {
		return controllers.NewVideoController(videoService)
	})
	if err != nil {
		panic(err)
	}
}
