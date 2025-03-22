package utils

// CheckMP4Header is a method to check MP4 header
func CheckMP4Header(header []byte) bool {
	if len(header) < 8 {
		return false
	}
	// check if the header starts with "ftyp"
	return string(header[4:8]) == "ftyp"
}
