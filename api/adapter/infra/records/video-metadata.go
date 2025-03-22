package records

const VideoMetadataTableName = "video_metadata"

// VideoMetadata is a record of video metadata
type VideoMetadata struct {
	// ID is a video ID
	ID string `dynamodbav:"id"`
	// Title is a title
	Title string `dynamodbav:"title"`
	// CreatedAt is a created at
	CreatedAt string `dynamodbav:"created_at"`
}
