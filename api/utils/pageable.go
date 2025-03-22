package utils

// Pageable is a struct of dynamodb pageable
type Pageable[T any] struct {
	// content is a content
	content []T
	// lastEvaluatedKey is a last evaluated key
	lastEvaluatedKey *string
}

// NewPageable is a constructor
func NewPageable[T any](content []T, lastEvaluatedKey *string) *Pageable[T] {
	return &Pageable[T]{
		content:          content,
		lastEvaluatedKey: lastEvaluatedKey,
	}
}

// LastEvaluatedKey is a getter
func (p *Pageable[T]) LastEvaluatedKey() *string {
	return p.lastEvaluatedKey
}

// Content is a getter
func (p *Pageable[T]) Content() []T {
	return p.content
}
