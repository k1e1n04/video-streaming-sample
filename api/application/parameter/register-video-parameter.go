package parameter

// RegisterVideoParameter is a parameter for RegisterVideo
type RegisterVideoParameter struct {
	// Title is a title
	Title string `json:"title"`
	// Video is a video
	Video []byte `json:"video"`
}
