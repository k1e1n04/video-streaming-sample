package utils

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

// GenerateID is a function to generate an ID
func GenerateID() string {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
