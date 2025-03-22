package entities

import (
	"fmt"

	"github.com/k1e1n04/video-streaming-sample/api/errors"
)

const videoTitleMaxLength = 100

// VideoTitle is a video title
type VideoTitle struct {
	// value is a value
	value string
}

// NewVideoTitle is a constructor
func NewVideoTitle(
	value string,
) (*VideoTitle, error) {
	if value == "" {
		return nil, errors.NewBadRequestError(
			"video title must not be empty",
			"video title must not be empty",
		)
	}
	if len(value) > videoTitleMaxLength {
		return nil, errors.NewBadRequestError(
			fmt.Sprintf("video title must be less than 100 characters: %s", value),
			"video title must be less than 100 characters",
		)
	}
	return &VideoTitle{
		value: value,
	}, nil
}

// Value is a getter
func (v *VideoTitle) Value() string {
	return v.value
}
