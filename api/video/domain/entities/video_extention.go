package entities

import "github.com/k1e1n04/video-streaming-sample/api/errors"

// VideoExtension is a video extension
type VideoExtension struct {
	// value is a value
	value string
}

// NewVideoExtension is a constructor
func NewVideoExtension(
	value string,
) (*VideoExtension, error) {
	if value == "" {
		return nil, errors.NewBadRequestError(
			"video extension must not be empty",
			"video extension must not be empty",
		)
	}
	if !isAllowedExtension(value) {
		return nil, errors.NewBadRequestError(
			"video extension must be mp4",
			"video extension must be mp4",
		)
	}
	return &VideoExtension{
		value: value,
	}, nil
}

// isAllowedExtension is a method to check if the extension is allowed
func isAllowedExtension(value string) bool {
	return value == "mp4"
}
