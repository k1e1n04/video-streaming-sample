package models

import "github.com/k1e1n04/video-streaming-sample/api/utils"

// UserID is a user ID
type UserID struct {
	// value is a user ID
	value string
}

// NewUserID is a constructor
func NewUserID() UserID {
	return UserID{
		value: utils.GenerateID(),
	}
}

// RestoreUserID is a constructor
func RestoreUserID(
	value string,
) UserID {
	return UserID{
		value: value,
	}
}

// Value is a getter
func (u *UserID) Value() string {
	return u.value
}

// Equals is a method to compare user IDs
func (u *UserID) Equals(target UserID) bool {
	return u.value == target.value
}
