package adapters

import (
	"context"
	"firebase.google.com/go/v4/db"
	"github.com/pkulik0/stredono/cloud/platform/modules"
)

type FirebaseRealtimeDbAdapter struct {
	Client *db.Client
}

type FirebaseRealtimeDbRefAdapter struct {
	Ref *db.Ref
}

func (a *FirebaseRealtimeDbAdapter) NewRef(path string) modules.Ref {
	return &FirebaseRealtimeDbRefAdapter{Ref: a.Client.NewRef(path)}
}

func (r *FirebaseRealtimeDbRefAdapter) Get(ctx context.Context, v interface{}) error {
	return r.Ref.Get(ctx, v)
}

func (r *FirebaseRealtimeDbRefAdapter) Set(ctx context.Context, v interface{}) error {
	return r.Ref.Set(ctx, v)
}

func (r *FirebaseRealtimeDbRefAdapter) Delete(ctx context.Context) error {
	return r.Ref.Delete(ctx)
}

func (r *FirebaseRealtimeDbRefAdapter) Child(path string) modules.Ref {
	return &FirebaseRealtimeDbRefAdapter{Ref: r.Ref.Child(path)}
}

func (r *FirebaseRealtimeDbRefAdapter) Push(ctx context.Context, v interface{}) (modules.Ref, error) {
	ref, err := r.Ref.Push(ctx, v)
	if err != nil {
		return nil, err
	}
	return &FirebaseRealtimeDbRefAdapter{Ref: ref}, nil
}

type FirebaseRealtimeDbTransactionNodeAdapter struct {
	Node db.TransactionNode
}

func (n *FirebaseRealtimeDbTransactionNodeAdapter) Unmarshal(v interface{}) error {
	return n.Node.Unmarshal(v)
}

func (r *FirebaseRealtimeDbRefAdapter) Transaction(ctx context.Context, f func(modules.TransactionNode) (interface{}, error)) error {
	return r.Ref.Transaction(ctx, func(node db.TransactionNode) (interface{}, error) {
		return f(&FirebaseRealtimeDbTransactionNodeAdapter{Node: node})
	})
}
