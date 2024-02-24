package adapters

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/pkulik0/stredono/cloud/platform/modules"
)

type FirestoreDatabase struct {
	Client *firestore.Client
}

func (fdb *FirestoreDatabase) Collection(path string) modules.Collection {
	return &FirestoreCollection{ref: fdb.Client.Collection(path)}
}

type FirestoreCollection struct {
	ref *firestore.CollectionRef
}

func (fc *FirestoreCollection) Doc(path string) modules.Document {
	return &FirestoreDocument{ref: fc.ref.Doc(path)}
}

func (fc *FirestoreCollection) Add(ctx context.Context, data interface{}) (*modules.AddResult, error) {
	docRef, writeResult, err := fc.ref.Add(ctx, data)
	if err != nil {
		return nil, err
	}

	return &modules.AddResult{
		WriteResult: modules.WriteResult{Time: writeResult.UpdateTime},
		Doc:         &FirestoreDocument{ref: docRef},
	}, nil
}

type FirestoreDocument struct {
	ref *firestore.DocumentRef
}

func (fd *FirestoreDocument) Id() string {
	return fd.ref.ID
}

func (fd *FirestoreDocument) Set(ctx context.Context, data interface{}, opts *modules.DbOpts) (*modules.WriteResult, error) {
	var fsOpts []firestore.SetOption
	if opts != nil {
		if opts.MergeAll {
			fsOpts = append(fsOpts, firestore.MergeAll)
		}
	}

	writeResult, err := fd.ref.Set(ctx, data, fsOpts...)
	return &modules.WriteResult{Time: writeResult.UpdateTime}, err
}

func (fd *FirestoreDocument) Create(ctx context.Context, data interface{}) (*modules.WriteResult, error) {
	writeResult, err := fd.ref.Create(ctx, data)
	return &modules.WriteResult{Time: writeResult.UpdateTime}, err
}

type FirestoreDocumentSnapshot struct {
	snapshot *firestore.DocumentSnapshot
}

func (fds FirestoreDocumentSnapshot) DataTo(v interface{}) error {
	return fds.snapshot.DataTo(v)
}

func (fd *FirestoreDocument) Get(ctx context.Context) (modules.DocumentSnapshot, error) {
	snapshot, err := fd.ref.Get(ctx)
	return FirestoreDocumentSnapshot{snapshot: snapshot}, err
}
