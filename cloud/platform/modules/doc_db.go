package modules

import (
	"context"
	"time"
)

type WriteResult struct {
	Time time.Time
}

type AddResult struct {
	WriteResult
	Doc Document
}

type DbOpts struct {
	MergeAll bool
}

type Document interface {
	Id() string
	Create(ctx context.Context, data interface{}) (*WriteResult, error)
	Set(ctx context.Context, data interface{}, opts DbOpts) (*WriteResult, error)
	Get(ctx context.Context) (DocumentSnapshot, error)
}

type DocumentSnapshot interface {
	DataTo(v interface{}) error
	Ref() Document
}

type Collection interface {
	Doc(path string) Document
	Add(ctx context.Context, data interface{}) (*AddResult, error)
	Where(field, op string, value interface{}) Query
}

type Query interface {
	Get(ctx context.Context) QueryIterator
	Where(field, op string, value interface{}) Query
}

type QueryIterator interface {
	All() ([]DocumentSnapshot, error)
	Next() (DocumentSnapshot, error)
	Stop()
}

type DocDb interface {
	Collection(path string) Collection
	RunTransaction(ctx context.Context, f func(ctx context.Context, tx Transaction) error) error
}

type Transaction interface {
	Get(doc Document) (DocumentSnapshot, error)
	Set(doc Document, data interface{}, opts DbOpts) error
	Documents(q Query) QueryIterator
}
