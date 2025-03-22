package parameter

// GetVideoPageParameter is a parameter for GetVideoPage
type GetVideoPageParameter struct {
	// LastEvaluatedKey is a last evaluated key
	LastEvaluatedKey *string
	// Limit is a limit
	Limit int32
}
