package utils

import "time"

const (
	// DefaultTimeFormat は 標準の日付フォーマット"2024-04-08 22:06:15"
	DefaultTimeFormat = "2006-01-02 15:04:05"
	// DefaultDateFormat は 標準の日付フォーマット
	DefaultDateFormat = "2006-01-02"
)

// GetNow は現在時刻を取得
func GetNow() time.Time {
	return time.Now().Round(0)
}

// ToDateTimeString は日時を文字列に変換
func ToDateTimeString(t time.Time) string {
	return t.Format(DefaultTimeFormat)
}

// ParseDateTime は文字列を日時に変換
func ParseDateTime(t string) (time.Time, error) {
	return time.Parse(DefaultTimeFormat, t)
}
