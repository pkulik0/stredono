package modules

import (
	"context"
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
	Name() string
}

type Writer interface {
	Write(p []byte) (n int, err error)
	Close() error
}
