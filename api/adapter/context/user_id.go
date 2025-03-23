package context2

import "context"

// UserIDKey は `context.Context` に保存する UserID のキー
type contextKey struct{}

var UserIDKey = &contextKey{}

// SetUserID は `UserID` をコンテキストに埋め込む
func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

// UserIDFromContext はコンテキストから `UserID` を取得する
func UserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(UserIDKey).(string)
	return userID, ok
}
