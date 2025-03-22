package entities

import (
	"fmt"

	"github.com/k1e1n04/video-streaming-sample/api/errors"
)

// VideoStatus は動画の状態を表す Enum
type VideoStatus int

const (
	VideoStatusPublic  VideoStatus = iota // is public
	VideoStatusPrivate                    // is private
	VideoStatusFailed                     // failed
)

// String is a getter of VideoStatus
func (s VideoStatus) String() string {
	switch s {
	case VideoStatusPublic:
		return "PUBLIC"
	case VideoStatusPrivate:
		return "PRIVATE"
	case VideoStatusFailed:
		return "FAILED"
	default:
		return "UNKNOWN"
	}
}

// NewVideoStatus is a constructor
func NewVideoStatus(
	s string,
) (VideoStatus, error) {
	switch s {
	case "PUBLIC":
		return VideoStatusPublic, nil
	case "PRIVATE":
		return VideoStatusPrivate, nil
	case "FAILED":
		return VideoStatusFailed, nil
	default:
		return VideoStatus(-1), errors.NewBadRequestError(
			fmt.Sprintf("invalid video status: %s", s),
			"invalid video status",
		)
	}
}

// RestoreVideoStatus is a parser of VideoStatus
func RestoreVideoStatus(s string) (VideoStatus, error) {
	switch s {
	case "PUBLIC":
		return VideoStatusPublic, nil
	case "PRIVATE":
		return VideoStatusPrivate, nil
	case "FAILED":
		return VideoStatusFailed, nil
	default:
		return VideoStatus(-1), errors.NewInvalidStatementError(
			fmt.Sprintf("invalid video status: %s", s),
			nil,
		)
	}
}
