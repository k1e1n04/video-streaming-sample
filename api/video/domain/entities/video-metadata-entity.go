package entities

import (
	"time"

	"github.com/k1e1n04/video-streaming-sample/api/shared/models"

	"github.com/k1e1n04/video-streaming-sample/api/errors"
	"github.com/k1e1n04/video-streaming-sample/api/utils"
)

// VideoMetadataEntity is a video metadata entity
type VideoMetadataEntity struct {
	// id is a video id
	id VideoID
	// videoExtension is a video extension
	videoExtension VideoExtension
	// userID is a user id
	userID models.UserID
	// title is a title
	title VideoTitle
	// description is a description
	description VideoDescription
	// thumbnailID is a thumbnail id
	thumbnailID ThumbnailID
	// thumbnailExtension is a thumbnail extension
	thumbnailExtension ThumbnailExtension
	// views is a views
	views int64
	// likes is a likes
	likes int64
	// duration is a duration
	duration int64
	// status is a status
	status VideoStatus
	// createdAt is a created at
	createdAt time.Time
}

// NewVideoMetadataEntity is a constructor
func NewVideoMetadataEntity(
	videoExtension string,
	title string,
	description string,
	thumbnailExtension string,
	duration int64,
	status VideoStatus,
	userID models.UserID,
) (*VideoMetadataEntity, error) {
	videoExt, err := NewVideoExtension(videoExtension)
	if err != nil {
		return nil, err
	}
	videoTitle, err := NewVideoTitle(title)
	if err != nil {
		return nil, err
	}
	videoDesc, err := NewVideoDescription(description)
	if err != nil {
		return nil, err
	}
	thumbnailID := NewThumbnailID()
	thumbnailExt, err := NewThumbnailExtension(thumbnailExtension)
	if err != nil {
		return nil, err
	}
	return &VideoMetadataEntity{
		id:                 NewVideoID(),
		videoExtension:     *videoExt,
		userID:             userID,
		title:              *videoTitle,
		description:        *videoDesc,
		thumbnailID:        thumbnailID,
		thumbnailExtension: *thumbnailExt,
		status:             status,
		views:              0,
		likes:              0,
		duration:           duration,
		createdAt:          utils.GetNow(),
	}, nil
}

// RestoreVideoMetadataEntity is a constructor
func RestoreVideoMetadataEntity(
	id string,
	userID string,
	videoExtension string,
	title string,
	thumbnailID string,
	thumbnailExtension string,
	description string,
	status string,
	likes int64,
	duration int64,
	views int64,
	createdAt string,
) (*VideoMetadataEntity, error) {
	videoID := RestoreVideoID(id)
	videoUserID := models.RestoreUserID(userID)
	videoExt, err := NewVideoExtension(videoExtension)
	if err != nil {
		return nil, errors.NewInvalidStatementError(
			err.Error(),
			err,
		)
	}
	videoTitle, err := NewVideoTitle(title)
	if err != nil {
		return nil, errors.NewInvalidStatementError(
			err.Error(),
			err,
		)
	}
	thID := RestoreThumbnailID(thumbnailID)
	thExt, err := NewThumbnailExtension(thumbnailExtension)
	if err != nil {
		return nil, errors.NewInvalidStatementError(
			err.Error(),
			err,
		)
	}
	videoDesc, err := NewVideoDescription(description)
	if err != nil {
		return nil, errors.NewInvalidStatementError(
			err.Error(),
			err,
		)
	}
	videoStatus, err := RestoreVideoStatus(status)
	if err != nil {
		return nil, err
	}
	createdAtTime, err := utils.ParseDateTime(createdAt)
	if err != nil {
		return nil, errors.NewInvalidStatementError(
			err.Error(),
			err,
		)
	}
	return &VideoMetadataEntity{
		id:                 videoID,
		userID:             videoUserID,
		videoExtension:     *videoExt,
		title:              *videoTitle,
		thumbnailID:        thID,
		thumbnailExtension: *thExt,
		description:        *videoDesc,
		status:             videoStatus,
		likes:              likes,
		views:              views,
		duration:           duration,
		createdAt:          createdAtTime,
	}, nil
}

// ID is a getter
func (v *VideoMetadataEntity) ID() *VideoID {
	return &v.id
}

// Title is a getter
func (v *VideoMetadataEntity) Title() *VideoTitle {
	return &v.title
}

// UserID is a getter
func (v *VideoMetadataEntity) UserID() models.UserID {
	return v.userID
}

// VideoExtension is a getter
func (v *VideoMetadataEntity) VideoExtension() *VideoExtension {
	return &v.videoExtension
}

// Description is a getter
func (v *VideoMetadataEntity) Description() *VideoDescription {
	return &v.description
}

// ThumbnailID is a getter
func (v *VideoMetadataEntity) ThumbnailID() ThumbnailID {
	return v.thumbnailID
}

// ThumbnailExtension is a getter
func (v *VideoMetadataEntity) ThumbnailExtension() *ThumbnailExtension {
	return &v.thumbnailExtension
}

// Views is a getter
func (v *VideoMetadataEntity) Views() int64 {
	return v.views
}

// Likes is a getter
func (v *VideoMetadataEntity) Likes() int64 {
	return v.likes
}

// Duration is a getter
func (v *VideoMetadataEntity) Duration() int64 {
	return v.duration
}

// Status is a getter
func (v *VideoMetadataEntity) Status() VideoStatus {
	return v.status
}

// CreatedAt is a getter
func (v *VideoMetadataEntity) CreatedAt() time.Time {
	return v.createdAt
}
