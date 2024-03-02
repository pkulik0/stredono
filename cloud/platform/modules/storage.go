package modules

import (
	"context"
	"time"
)

type Storage interface {
	DefaultBucket() (Bucket, error)
	Bucket(name string) (Bucket, error)
}

type Bucket interface {
	Object(name string) Object
}

type Object interface {
	NewWriter(ctx context.Context) Writer
	Attrs(ctx context.Context) (*ObjectAttrs, error)
	SetPublicRead(ctx context.Context) error
	Name() string
}

type ObjectAttrs struct {
	Name        string
	ContentType string
	Size        int64
	CreatedAt   time.Time
	MediaUrl    string
}

type Writer interface {
	Write(p []byte) (n int, err error)
	Close() error
}
