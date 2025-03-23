package controllers

import (
	"bytes"
	"context"
	context2 "github.com/k1e1n04/video-streaming-sample/api/adapter/context"
	"github.com/k1e1n04/video-streaming-sample/api/video/domain/entities"
	"io"
	"log"

	parameter2 "github.com/k1e1n04/video-streaming-sample/api/video/application/parameter"
	"github.com/k1e1n04/video-streaming-sample/api/video/application/services"

	"github.com/k1e1n04/video-streaming-sample/api/adapter/grpc/video"
	"github.com/k1e1n04/video-streaming-sample/api/errors"
	"github.com/k1e1n04/video-streaming-sample/api/utils"
	"google.golang.org/grpc"
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

// UploadVideo is a method to upload video (MP4 only)
func (v *VideoController) UploadVideo(server grpc.ClientStreamingServer[video.UploadVideoRequest, video.UploadVideoResponse]) error {
	var metadata *video.VideoMetadata
	var videoBuffer bytes.Buffer
	var thumbnailBuffer bytes.Buffer
	var headerBuffer bytes.Buffer
	isMP4 := false
	receivedBytes := 0
	ctx := server.Context()

	userID, ok := context2.UserIDFromContext(ctx)
	if !ok {
		return errors.NewUnauthorizedError(
			"failed to get user ID",
			"unauthorized",
			nil,
		)
	}

	for {
		req, err := server.Recv()
		if err == io.EOF {
			// check if the video file is empty
			if videoBuffer.Len() == 0 {
				return errors.NewBadRequestError(
					"video file must not be empty",
					"video file must not be empty",
				)
			}

			if !isMP4 {
				return errors.NewBadRequestError(
					"video file must be MP4",
					"video file must be MP4",
				)
			}

			status, err := entities.NewVideoStatus(int(metadata.GetStatus().Number()))
			if err != nil {
				return err
			}
			id, err := v.videoService.Register(ctx, parameter2.RegisterVideoParameter{
				UserID:             userID,
				Title:              metadata.Title,
				Description:        metadata.Description,
				Extension:          metadata.Extension,
				Duration:           metadata.Duration,
				Status:             status,
				ThumbnailExtension: metadata.ThumbnailExtension,
				Thumbnail:          bytes.NewReader(thumbnailBuffer.Bytes()),
				Video:              bytes.NewReader(videoBuffer.Bytes()),
			})
			if err != nil {
				return err
			}

			return server.SendAndClose(&video.UploadVideoResponse{VideoId: *id})
		}
		if err != nil {
			return errors.NewInvalidStatementError(
				"failed to receive a request",
				err,
			)
		}

		switch req.Data.(type) {
		case *video.UploadVideoRequest_Metadata:
			metadata = req.GetMetadata()

		case *video.UploadVideoRequest_Chunk:
			data := req.GetChunk()

			if len(data) == 0 {
				log.Printf("Warning: Received an empty chunk")
			}

			// check if the video file is MP4
			if !isMP4 && headerBuffer.Len() < 8 {
				headerBuffer.Write(data)

				if headerBuffer.Len() >= 8 {
					if utils.CheckMP4Header(headerBuffer.Bytes()) {
						isMP4 = true
					}
				}
			}

			receivedBytes += len(data)
			videoBuffer.Write(data)
		case *video.UploadVideoRequest_Thumbnail:
			thumbnailBuffer.Write(req.GetThumbnail())
		}
	}
}

// GetVideoURL is a method to get video URL
func (v *VideoController) GetVideoURL(ctx context.Context, req *video.GetVideoRequest) (*video.GetVideoResponse, error) {
	_, ok := context2.UserIDFromContext(ctx)
	if !ok {
		return nil, errors.NewUnauthorizedError(
			"failed to get user ID",
			"unauthorized",
			nil,
		)
	}
	url, err := v.videoService.GetPresignedURLByVideoID(ctx, parameter2.GetPresignedURLParameter{
		VideoID: req.VideoId,
	})
	if err != nil {
		return nil, err
	}
	return &video.GetVideoResponse{
		PresignedUrl: *url,
	}, nil
}

// ListVideos is a method to list videos
func (v *VideoController) ListVideos(ctx context.Context, req *video.ListVideosRequest) (*video.ListVideosResponse, error) {
	_, ok := context2.UserIDFromContext(ctx)
	if !ok {
		return nil, errors.NewUnauthorizedError(
			"failed to get user ID",
			"unauthorized",
			nil,
		)
	}
	videoPage, err := v.videoService.GetVideoPage(ctx, parameter2.GetVideoPageParameter{
		Limit:            req.Limit,
		LastEvaluatedKey: req.LastEvaluatedKey,
	})
	if err != nil {
		return nil, err
	}

	res := make([]*video.VideoInfo, 0, len(videoPage.Content()))
	for _, videoMetadata := range videoPage.Content() {
		res = append(res, &video.VideoInfo{
			VideoId:   videoMetadata.ID,
			Title:     videoMetadata.Title,
			CreatedAt: utils.ToDateTimeString(videoMetadata.CreatedAt),
		})
	}

	return &video.ListVideosResponse{
		Videos:           res,
		LastEvaluatedKey: videoPage.LastEvaluatedKey(),
	}, nil
}
