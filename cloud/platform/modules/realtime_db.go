package modules

import "context"

type RealtimeDb interface {
	NewRef(path string) Ref
}

type Ref interface {
	Get(ctx context.Context, v interface{}) error
	Set(ctx context.Context, v interface{}) error
	Delete(ctx context.Context) error
	Child(path string) Ref
}
