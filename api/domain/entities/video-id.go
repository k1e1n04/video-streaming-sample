package entities

import "github.com/k1e1n04/video-streaming-sample/api/utils"

type VideoID struct {
	// value is a value
	value string
}

// NewVideoID is a constructor
func NewVideoID() VideoID {
	return VideoID{
		value: utils.GenerateID(),
	}
}

// RestoreVideoID is a constructor
func RestoreVideoID(
	value string,
) VideoID {
	return VideoID{
		value: value,
	}
}

// Value is a getter
func (v *VideoID) Value() string {
	return v.value
}
