package controllers

import (
	"bytes"
	"context"
	"github.com/k1e1n04/video-streaming-sample/api/adapter/grpc/video"
	"github.com/k1e1n04/video-streaming-sample/api/application/parameter"
	"github.com/k1e1n04/video-streaming-sample/api/application/services"
	"github.com/k1e1n04/video-streaming-sample/api/errors"
	"github.com/k1e1n04/video-streaming-sample/api/utils"
	"google.golang.org/grpc"
	"io"
	"log"
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
	var title string
	var videoBuffer bytes.Buffer
	var headerBuffer bytes.Buffer
	isMP4 := false
	receivedBytes := 0

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

			id, err := v.videoService.Register(context.Background(), parameter.RegisterVideoParameter{
				Title: title,
				Video: bytes.NewReader(videoBuffer.Bytes()),
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
		case *video.UploadVideoRequest_Title:
			title = req.GetTitle()

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
		}
	}
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
