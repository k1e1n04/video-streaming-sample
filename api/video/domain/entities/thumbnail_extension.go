package entities

import "github.com/k1e1n04/video-streaming-sample/api/errors"

// ThumbnailExtension is a thumbnail extension
type ThumbnailExtension struct {
	// value is a value
	value string
}

// NewThumbnailExtension is a constructor
func NewThumbnailExtension(
	value string,
) (*ThumbnailExtension, error) {
	if value == "" {
		return nil, errors.NewBadRequestError(
			"thumbnail extension must not be empty",
			"thumbnail extension must not be empty",
		)
	}
	if !isAllowedThumbnailExtension(value) {
		return nil, errors.NewBadRequestError(
			"thumbnail extension must be jpg, jpeg, or png",
			"thumbnail extension must be jpg, jpeg, or png",
		)
	}
	return &ThumbnailExtension{
		value: value,
	}, nil
}

func isAllowedThumbnailExtension(value string) bool {
	return value == "jpg" || value == "jpeg" || value == "png"
}

// Value is a getter
func (t *ThumbnailExtension) Value() string {
	return t.value
}
