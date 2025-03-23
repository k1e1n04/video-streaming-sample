package records

const VideoMetadataTableName = "video_metadata"

// VideoMetadata is a record of video metadata
type VideoMetadata struct {
	// ID is a video ID
	ID string `dynamodbav:"id"`
	// VideoExtension is a video extension
	VideoExtension string `dynamodbav:"video_extension"`
	// UserID is a user ID
	UserID string `dynamodbav:"user_id"`
	// Title is a title
	Title string `dynamodbav:"title"`
	// Description is a description
	Description string `dynamodbav:"description"`
	// ThumbnailID is a thumbnail ID
	ThumbnailID string `dynamodbav:"thumbnail_id"`
	// ThumbnailExtension is a thumbnail extension
	ThumbnailExtension string `dynamodbav:"thumbnail_extension"`
	// Views is a views
	Views int64 `dynamodbav:"views"`
	// Likes is a likes
	Likes int64 `dynamodbav:"likes"`
	// Duration is a duration
	Duration int64 `dynamodbav:"duration"`
	// Status is a status
	Status string `dynamodbav:"status"`
	// CreatedAt is a created at
	CreatedAt string `dynamodbav:"created_at"`
}
