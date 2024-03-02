package adapters

import (
	gcStorage "cloud.google.com/go/storage"
	"context"
	"errors"
	"firebase.google.com/go/v4/storage"
	"github.com/pkulik0/stredono/cloud/platform"
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
		if errors.Is(err, gcStorage.ErrBucketNotExist) {
			return nil, platform.ErrorBucketNotFound
		}
		return nil, err
	}
	return &GoogleBucket{Bucket: bucket}, nil
}

func (fs *FirebaseStorage) DefaultBucket() (modules.Bucket, error) {
	bucket, err := fs.Client.DefaultBucket()
	if err != nil {
		if errors.Is(err, gcStorage.ErrBucketNotExist) {
			return nil, platform.ErrorBucketNotFound
		}
		return nil, err
	}
	return &GoogleBucket{Bucket: bucket}, nil
}

func (g *GoogleBucket) Object(name string) modules.Object {
	obj := g.Bucket.Object(name)
	return &GoogleObject{Object: obj}
}

func (g *GoogleObject) SetPublicRead(ctx context.Context) error {
	return g.Object.ACL().Set(ctx, gcStorage.AllUsers, gcStorage.RoleReader)
}

func (g *GoogleObject) Attrs(ctx context.Context) (*modules.ObjectAttrs, error) {
	attrs, err := g.Object.Attrs(ctx)
	if err != nil {
		if errors.Is(err, gcStorage.ErrObjectNotExist) {
			return nil, platform.ErrorObjectNotFound
		}
		return nil, err
	}
	return &modules.ObjectAttrs{
		Name:        attrs.Name,
		ContentType: attrs.ContentType,
		Size:        attrs.Size,
		CreatedAt:   attrs.Created,
		MediaUrl:    attrs.MediaLink,
	}, nil
}

func (g *GoogleObject) NewWriter(ctx context.Context) modules.Writer {
	return g.Object.NewWriter(ctx)
}

func (g *GoogleObject) Name() string {
	return g.Object.ObjectName()
}
