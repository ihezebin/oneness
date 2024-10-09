package oss

import (
	"bytes"
	"context"
	"io"
	"net/url"
	"time"

	"github.com/pkg/errors"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

type kodoClient struct {
	uploader      *storage.FormUploader
	bucketManager *storage.BucketManager
	bucket        string
	accessKey     string
	secretKey     string
}

func (c *kodoClient) StatObject(ctx context.Context, name string) (*ObjectInfo, error) {
	info, err := c.bucketManager.Get(c.bucket, name, &storage.GetObjectInput{
		PresignUrl: true, // 下载 URL 是否进行签名，源站域名或者私有空间需要配置为 true
	})
	if err != nil {
		return nil, errors.Wrapf(err, "minio stat object err")
	}

	return &ObjectInfo{
		ETag:         info.ETag,
		Size:         info.ContentLength,
		ContentType:  info.ContentType,
		LastModified: info.LastModified,
	}, nil
}

func (c *kodoClient) GetObjects(ctx context.Context, prefix string, opts ...GetObjectsOption) ([]Object, error) {
	return nil, errors.New("not implement")
}

func (c *kodoClient) GetObject(ctx context.Context, name string) (io.ReadCloser, error) {
	obj, err := c.bucketManager.Get(c.bucket, name, &storage.GetObjectInput{
		PresignUrl: true, // 下载 URL 是否进行签名，源站域名或者私有空间需要配置为 true
	})
	if err != nil {
		return nil, errors.Wrapf(err, "kodo get object err")
	}

	body, err := io.ReadAll(obj.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "kodo read object err")
	}

	reader := io.NopCloser(bytes.NewReader(body))

	return reader, nil
}

func (c *kodoClient) PutObject(ctx context.Context, name string, reader io.Reader, opts ...PutOption) error {
	data, err := io.ReadAll(reader)
	if err != nil {
		return errors.Wrapf(err, "kodo read object err")
	}

	opts = append([]PutOption{WithSize(int64(len(data)))}, opts...)
	opt := newPutOptions(opts...)

	kodoOpt := &storage.PutExtra{}
	if opt.ContentType != "" {
		kodoOpt.MimeType = opt.ContentType
	}

	ret := storage.PutRet{}

	if err = c.uploader.Put(ctx, &ret, c.token(), name, bytes.NewReader(data), opt.Size, kodoOpt); err != nil {
		return errors.Wrapf(err, "kodo put object err")
	}

	return nil
}

func (c *kodoClient) RemoveObject(ctx context.Context, name string) error {
	if err := c.bucketManager.Delete(c.bucket, name); err != nil {
		return errors.Wrapf(err, "kodo remove object err")
	}
	return nil
}

func (c *kodoClient) SignedURL(ctx context.Context, name string, method string, expires time.Duration) (*url.URL, error) {
	deadline := time.Now().Add(expires).Unix() //1小时有效期
	privateAccessURL := storage.MakePrivateURL(auth.New(c.accessKey, c.secretKey), "", name, deadline)

	u, err := url.Parse(privateAccessURL)
	if err != nil {
		return nil, errors.Wrapf(err, "kodo parse url err")
	}

	return u, nil
}

func (c *kodoClient) token() string {
	putPolicy := storage.PutPolicy{
		Scope: c.bucket,
	}
	token := putPolicy.UploadToken(auth.New(c.accessKey, c.secretKey))
	return token
}

var _ Client = (*kodoClient)(nil)

// NewKodoClient
// endpoint 值为: ["z0", "cn-east-2", "z1", "z2", "na0", "as0"]
//
//	RIDHuadong          = RegionID("z0")
//	RIDHuadongZheJiang2 = RegionID("cn-east-2")
//	RIDHuabei           = RegionID("z1")
//	RIDHuanan           = RegionID("z2")
//	RIDNorthAmerica     = RegionID("na0")
//	RIDSingapore        = RegionID("as0")
func NewKodoClient(endpoint string, accessKey string, secretKey string, bucket string) (*kodoClient, error) {
	if bucket == "" {
		return nil, errors.New("minio bucket is empty")
	}

	regionId := storage.RegionID(endpoint)
	//region, ok := storage.GetRegionByID(regionId)
	//if !ok {
	//	return nil, errors.New("kodo invalid endpoint")
	//}

	config := &storage.Config{
		Region:   &storage.ZoneHuadong,
		UseHTTPS: true,
	}
	uploader := storage.NewFormUploader(config)

	bucketManager := storage.NewBucketManager(auth.New(accessKey, secretKey), config)

	err := bucketManager.CreateBucket(bucket, regionId)
	if err != nil {
		return nil, errors.Wrapf(err, "kodo create bucket err")
	}

	return &kodoClient{
		uploader:      uploader,
		bucketManager: bucketManager,
		bucket:        bucket,
	}, nil
}

func (c *kodoClient) BucketManager() *storage.BucketManager {
	return c.bucketManager
}
func (c *kodoClient) Uploader() *storage.FormUploader {
	return c.uploader
}
