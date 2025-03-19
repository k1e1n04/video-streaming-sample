package entities

import (
	"github.com/k1e1n04/video-streaming-sample/api/errors"
	"github.com/k1e1n04/video-streaming-sample/api/utils"
	"time"
)

// VideoMetadataEntity is a video metadata entity
type VideoMetadataEntity struct {
	// id is a video id
	id VideoID
	// title is a title
	title VideoTitle
	// createdAt is a created at
	createdAt time.Time
}

// NewVideoMetadataEntity is a constructor
func NewVideoMetadataEntity(
	title string,
) (*VideoMetadataEntity, error) {
	videoTitle, err := NewVideoTitle(title)
	if err != nil {
		return nil, err
	}
	return &VideoMetadataEntity{
		id:        NewVideoID(),
		title:     *videoTitle,
		createdAt: utils.GetNow(),
	}, nil
}

// RestoreVideoMetadataEntity is a constructor
func RestoreVideoMetadataEntity(
	id string,
	title string,
	createdAt string,
) (*VideoMetadataEntity, error) {
	videoID := RestoreVideoID(id)
	videoTitle, err := NewVideoTitle(title)
	if err != nil {
		return nil, errors.NewInvalidStatementError(
			err.Error(),
			err,
		)
	}
	createdAtTime, err := utils.ParseDateTime(createdAt)
	if err != nil {
		return nil, errors.NewInvalidStatementError(
			err.Error(),
			err,
		)
	}
	return &VideoMetadataEntity{
		id:        videoID,
		title:     *videoTitle,
		createdAt: createdAtTime,
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

// CreatedAt is a getter
func (v *VideoMetadataEntity) CreatedAt() time.Time {
	return v.createdAt
}
