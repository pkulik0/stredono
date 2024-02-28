package adapters

import (
	gcStorage "cloud.google.com/go/storage"
	"context"
	"firebase.google.com/go/v4/storage"
	"github.com/pkulik0/stredono/cloud/platform/modules"
)

type FirebaseStorage struct {
	Client *storage.Client
}

type GoogleBucket struct {
	Bucket *gcStorage.BucketHandle
}

type GoogleObject struct {
	Object *gcStorage.ObjectHandle
}

func (fs *FirebaseStorage) Bucket(name string) (modules.Bucket, error) {
	bucket, err := fs.Client.Bucket(name)
	if err != nil {
		return nil, err
	}
	return &GoogleBucket{Bucket: bucket}, nil
}

func (fs *FirebaseStorage) DefaultBucket() (modules.Bucket, error) {
	bucket, err := fs.Client.DefaultBucket()
	if err != nil {
		return nil, err
	}
	return &GoogleBucket{Bucket: bucket}, nil
}

func (g *GoogleBucket) Object(name string) modules.Object {
	obj := g.Bucket.Object(name)
	return &GoogleObject{Object: obj}
}

func (g *GoogleObject) NewWriter(ctx context.Context) modules.Writer {
	return g.Object.NewWriter(ctx)
}

func (g *GoogleObject) Name() string {
	return g.Object.ObjectName()
}
