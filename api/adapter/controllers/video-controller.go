package controllers

import (
	"context"
	"github.com/k1e1n04/video-streaming-sample/api/adapter/grpc/video"
	"github.com/k1e1n04/video-streaming-sample/api/application/parameter"
	"github.com/k1e1n04/video-streaming-sample/api/application/services"
)

// VideoController is a video controller
type VideoController struct {
	video.UnimplementedVideoServiceServer
	videoService services.VideoService
}

// NewVideoController is a constructor
func NewVideoController(videoService services.VideoService) VideoController {
	return VideoController{
		videoService: videoService,
	}
}

// UploadVideo is a method to upload video
func (v *VideoController) UploadVideo(ctx context.Context, req *video.UploadVideoRequest) (*video.UploadVideoResponse, error) {
	id, err := v.videoService.Register(ctx, parameter.RegisterVideoParameter{
		Title: req.Title,
		Video: req.File,
	})
	if err != nil {
		return nil, err
	}
	return &video.UploadVideoResponse{
		VideoId: *id,
	}, nil
}

// GetVideoURL is a method to get video URL
func (v *VideoController) GetVideoURL(ctx context.Context, req *video.GetVideoRequest) (*video.GetVideoResponse, error) {
	url, err := v.videoService.GetPresignedURLByVideoID(ctx, parameter.GetPresignedURLParameter{
		VideoID: req.VideoId,
	})
	if err != nil {
		return nil, err
	}
	return &video.GetVideoResponse{
		PresignedUrl: *url,
	}, nil
}
