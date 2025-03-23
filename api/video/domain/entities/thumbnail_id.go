package entities

import "github.com/k1e1n04/video-streaming-sample/api/utils"

// ThumbnailID is a thumbnail ID
type ThumbnailID struct {
	// value is a value
	value string
}

// NewThumbnailID is a constructor
func NewThumbnailID() ThumbnailID {
	return ThumbnailID{
		value: utils.GenerateID(),
	}
}

// RestoreThumbnailID is a constructor
func RestoreThumbnailID(
	value string,
) ThumbnailID {
	return ThumbnailID{
		value: value,
	}
}

// Value is a getter
func (t *ThumbnailID) Value() string {
	return t.value
}

// Equals is a method to compare thumbnail IDs
func (t *ThumbnailID) Equals(target ThumbnailID) bool {
	return t.value == target.value
}
