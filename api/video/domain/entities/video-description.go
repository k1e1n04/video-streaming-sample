package entities

import (
	"fmt"
	"github.com/k1e1n04/video-streaming-sample/api/errors"
)

// VideoDescription is a video description
type VideoDescription struct {
	// value is a value
	value string
}

const videoDescriptionMaxLength = 100

// NewVideoDescription is a constructor
func NewVideoDescription(
	value string,
) (*VideoDescription, error) {
	if len(value) > videoDescriptionMaxLength {
		return nil, errors.NewBadRequestError(
			fmt.Sprintf("video description must be less than 100 characters: %s", value),
			"video description must be less than 100 characters",
		)
	}
	return &VideoDescription{
		value: value,
	}, nil
}

// Value is a getter
func (v *VideoDescription) Value() string {
	return v.value
}
