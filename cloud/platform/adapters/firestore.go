package adapters

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/pkulik0/stredono/cloud/platform"
)

type FirestoreCollection struct {
	ref *firestore.CollectionRef
}

func (fc *FirestoreCollection) Doc(path string) platform.Document {
	return &FirestoreDocument{ref: fc.ref.Doc(path)}
}

type FirestoreDocument struct {
	ref *firestore.DocumentRef
}

func (fd *FirestoreDocument) Set(ctx context.Context, data interface{}, opts *platform.DbOpts) (*platform.WriteResult, error) {
	var fsOpts []firestore.SetOption
	if opts != nil {
		if opts.MergeAll {
			fsOpts = append(fsOpts, firestore.MergeAll)
		}
	}

	writeResult, err := fd.ref.Set(ctx, data, fsOpts...)
	return &platform.WriteResult{Time: writeResult.UpdateTime}, err
}

func (fd *FirestoreDocument) Create(ctx context.Context, data interface{}) (*platform.WriteResult, error) {
	writeResult, err := fd.ref.Create(ctx, data)
	return &platform.WriteResult{Time: writeResult.UpdateTime}, err
}

type FirestoreDocumentSnapshot struct {
	snapshot *firestore.DocumentSnapshot
}

func (fds FirestoreDocumentSnapshot) DataTo(v interface{}) error {
	return fds.snapshot.DataTo(v)
}

func (fd *FirestoreDocument) Get(ctx context.Context) (platform.DocumentSnapshot, error) {
	snapshot, err := fd.ref.Get(ctx)
	return FirestoreDocumentSnapshot{snapshot: snapshot}, err
}
