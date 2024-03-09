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
