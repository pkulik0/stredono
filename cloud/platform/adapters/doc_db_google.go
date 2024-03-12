package adapters

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/pkulik0/stredono/cloud/platform/modules"
)

func dbOptsToFirebase(opts modules.DbOpts) []firestore.SetOption {
	var fsOpts []firestore.SetOption
	if opts.MergeAll {
		fsOpts = append(fsOpts, firestore.MergeAll)
	}

	return fsOpts
}

type FirestoreDatabase struct {
	Client *firestore.Client
}

func (fdb *FirestoreDatabase) Collection(path string) modules.Collection {
	return &FirestoreCollection{ref: fdb.Client.Collection(path)}
}

func (fdb *FirestoreDatabase) RunTransaction(ctx context.Context, f func(ctx context.Context, tx modules.Transaction) error) error {
	return fdb.Client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		return f(ctx, &FirestoreTransaction{tx: tx})
	})
}

type FirestoreTransaction struct {
	tx *firestore.Transaction
}

func (ft *FirestoreTransaction) Get(doc modules.Document) (modules.DocumentSnapshot, error) {
	snapshot, err := ft.tx.Get(doc.(*FirestoreDocument).ref)
	return FirestoreDocumentSnapshot{snapshot: snapshot}, err
}

func (ft *FirestoreTransaction) Set(doc modules.Document, data interface{}, opts modules.DbOpts) error {
	return ft.tx.Set(doc.(*FirestoreDocument).ref, data, dbOptsToFirebase(opts)...)
}

func (ft *FirestoreTransaction) Documents(q modules.Query) modules.QueryIterator {
	return &FirestoreQueryIterator{iterator: ft.tx.Documents(q.(*FirestoreQuery).query)}
}

type FirestoreCollection struct {
	ref *firestore.CollectionRef
}

func (fc *FirestoreCollection) Where(field string, op string, value interface{}) modules.Query {
	q := fc.ref.Where(field, op, value)
	return &FirestoreQuery{query: &q}
}

type FirestoreQuery struct {
	query *firestore.Query
}

func (fq *FirestoreQuery) Get(ctx context.Context) modules.QueryIterator {
	return &FirestoreQueryIterator{iterator: fq.query.Documents(ctx)}
}

func (fq *FirestoreQuery) Where(field string, op string, value interface{}) modules.Query {
	q := fq.query.Where(field, op, value)
	return &FirestoreQuery{query: &q}
}

func (fq *FirestoreQuery) Limit(n int) modules.Query {
	q := fq.query.Limit(n)
	return &FirestoreQuery{query: &q}
}

type FirestoreQueryIterator struct {
	iterator *firestore.DocumentIterator
}

func (fqi *FirestoreQueryIterator) All() ([]modules.DocumentSnapshot, error) {
	all, err := fqi.iterator.GetAll()
	if err != nil {
		return nil, err
	}

	var result []modules.DocumentSnapshot
	for _, s := range all {
		result = append(result, FirestoreDocumentSnapshot{snapshot: s})
	}
	return result, nil
}

func (fqi *FirestoreQueryIterator) Next() (modules.DocumentSnapshot, error) {
	snapshot, err := fqi.iterator.Next()
	if err != nil {
		return nil, err
	}
	return FirestoreDocumentSnapshot{snapshot: snapshot}, nil
}

func (fqi *FirestoreQueryIterator) Stop() {
	fqi.iterator.Stop()
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

func (fd *FirestoreDocument) Set(ctx context.Context, data interface{}, opts modules.DbOpts) (*modules.WriteResult, error) {
	writeResult, err := fd.ref.Set(ctx, data, dbOptsToFirebase(opts)...)
	if err != nil {
		return nil, err
	}
	return &modules.WriteResult{Time: writeResult.UpdateTime}, nil
}

func (fd *FirestoreDocument) Create(ctx context.Context, data interface{}) (*modules.WriteResult, error) {
	writeResult, err := fd.ref.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &modules.WriteResult{Time: writeResult.UpdateTime}, nil
}

func (fd *FirestoreDocument) Update(ctx context.Context, updates []modules.Update) (*modules.WriteResult, error) {
	var fsUpdates []firestore.Update
	for _, u := range updates {
		fsUpdates = append(fsUpdates, firestore.Update{Path: u.Path, Value: u.Value})
	}

	writeResult, err := fd.ref.Update(ctx, fsUpdates)
	if err != nil {
		return nil, err
	}
	return &modules.WriteResult{Time: writeResult.UpdateTime}, nil
}

type FirestoreDocumentSnapshot struct {
	snapshot *firestore.DocumentSnapshot
}

func (fds FirestoreDocumentSnapshot) Ref() modules.Document {
	return &FirestoreDocument{ref: fds.snapshot.Ref}
}

func (fds FirestoreDocumentSnapshot) DataTo(v interface{}) error {
	return fds.snapshot.DataTo(v)
}

func (fd *FirestoreDocument) Get(ctx context.Context) (modules.DocumentSnapshot, error) {
	snapshot, err := fd.ref.Get(ctx)
	if err != nil {
		return nil, err
	}
	return FirestoreDocumentSnapshot{snapshot: snapshot}, nil
}
