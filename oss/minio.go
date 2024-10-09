package oss

import (
	"context"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
)

type minioClient struct {
	kernel *minio.Client
	bucket string
}

func (c *minioClient) StatObject(ctx context.Context, name string) (*ObjectInfo, error) {
	info, err := c.kernel.StatObject(ctx, c.bucket, name, minio.StatObjectOptions{})
	if err != nil {
		return nil, errors.Wrapf(err, "minio stat object err")
	}

	return &ObjectInfo{
		Key:          info.Key,
		ETag:         info.ETag,
		Size:         info.Size,
		ContentType:  info.ContentType,
		LastModified: info.LastModified,
		Expires:      info.Expires,
	}, nil
}

func (c *minioClient) GetObjects(ctx context.Context, prefix string, opts ...GetObjectsOption) ([]Object, error) {
	objectsCh := c.kernel.ListObjects(ctx, c.bucket, minio.ListObjectsOptions{Prefix: prefix})

	objs := make([]Object, 0)

	for obj := range objectsCh {

		object, err := c.kernel.GetObject(ctx, c.bucket, obj.Key, minio.GetObjectOptions{})
		if err != nil {
			return nil, errors.Wrapf(err, "minio get object err")
		}

		objs = append(objs, Object{
			Key:  obj.Key,
			Data: object,
		})
	}

	return objs, nil
}

func (c *minioClient) GetObject(ctx context.Context, name string) (io.ReadCloser, error) {
	obj, err := c.kernel.GetObject(ctx, c.bucket, name, minio.GetObjectOptions{})
	if err != nil {
		return nil, errors.Wrapf(err, "minio get object err")
	}
	if _, err = obj.Stat(); err != nil {
		_ = obj.Close()
		return nil, errors.Wrapf(err, "minio get object stat err")
	}

	return obj, nil
}

func (c *minioClient) PutObject(ctx context.Context, name string, reader io.Reader, opts ...PutOption) error {
	opt := newPutOptions(opts...)

	miniOpt := minio.PutObjectOptions{}
	if opt.ContentType != "" {
		miniOpt.ContentType = opt.ContentType
	}

	if _, err := c.kernel.PutObject(ctx, c.bucket, name, reader, opt.Size, miniOpt); err != nil {
		return errors.Wrapf(err, "minio put object err")
	}

	return nil
}

func (c *minioClient) RemoveObject(ctx context.Context, name string) error {
	if err := c.kernel.RemoveObject(ctx, c.bucket, name, minio.RemoveObjectOptions{}); err != nil {
		return errors.Wrapf(err, "minio remove object err")
	}
	return nil
}

func (c *minioClient) SignedURL(ctx context.Context, name string, method string, expires time.Duration) (*url.URL, error) {
	u, err := c.kernel.Presign(ctx, method, c.bucket, name, expires, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "minio signed url err")
	}

	return u, nil
}

var _ Client = (*minioClient)(nil)

func NewMinioClient(endpoint string, accessKey string, secretKey string, bucket string) (*minioClient, error) {
	if bucket == "" {
		return nil, errors.New("minio bucket is empty")
	}

	client, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKey, secretKey, ""),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "new minio client err")
	}

	ctx := context.Background()
	exists, err := client.BucketExists(ctx, bucket)
	if err != nil {
		return nil, errors.Wrapf(err, "minio bucket exists err")
	}
	if !exists {
		if err = client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{}); err != nil {
			return nil, errors.Wrapf(err, "make minio bucket err")
		}
	}

	return &minioClient{
		kernel: client,
		bucket: bucket,
	}, nil
}

func (c *minioClient) Kernel() *minio.Client {
	return c.kernel
}
