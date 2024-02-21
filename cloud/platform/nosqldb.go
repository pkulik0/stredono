package platform

import (
	"context"
	"time"
)

type WriteResult struct {
	Time time.Time
}

type DbOpts struct {
	MergeAll bool
}

type Document interface {
	Create(ctx context.Context, data interface{}) (*WriteResult, error)
	Set(ctx context.Context, data interface{}, opts *DbOpts) (*WriteResult, error)
	Get(ctx context.Context) (DocumentSnapshot, error)
}

type DocumentSnapshot interface {
	DataTo(v interface{}) error
}

type Collection interface {
	Doc(path string) Document
}

type NoSqlDb interface {
	Collection(path string) Collection
}
