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
	Transaction(ctx context.Context, f func(TransactionNode) (interface{}, error)) error
	Push(ctx context.Context, v interface{}) (Ref, error)
}

type TransactionNode interface {
	Unmarshal(v interface{}) error
}
