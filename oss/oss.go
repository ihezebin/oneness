package oss

import (
	"context"
	"io"
	"net/url"
	"strconv"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg/errors"
)

type ossClient struct {
	kernel *oss.Client
	bucket *oss.Bucket
}

func (c *ossClient) StatObject(ctx context.Context, name string) (*ObjectInfo, error) {
	header, err := c.bucket.GetObjectDetailedMeta(name)
	if err != nil {
		return nil, errors.Wrapf(err, "oss get object detailed meta err")

	}

	info := &ObjectInfo{
		ContentType: header.Get("Content-Type"),
		ETag:        header.Get("Etag"),
	}
	if v := header.Get("Content-Length"); v != "" {
		info.Size, _ = strconv.ParseInt(v, 10, 64)
	}
	if v := header.Get("Last-Modified"); v != "" {
		info.LastModified, _ = time.Parse(time.RFC1123, v)
	}

	return info, nil
}

func (c *ossClient) GetObjects(ctx context.Context, prefix string, opts ...GetObjectsOption) ([]Object, error) {
	return nil, errors.New("not implement")
}

func (c *ossClient) GetObject(ctx context.Context, name string) (io.ReadCloser, error) {
	obj, err := c.bucket.GetObject(name)
	if err != nil {
		return nil, errors.Wrapf(err, "oss get object err")
	}

	return obj, nil
}

func (c *ossClient) PutObject(ctx context.Context, name string, reader io.Reader, opts ...PutOption) error {
	opt := newPutOptions(opts...)
	ossOpts := make([]oss.Option, 0)
	if opt.Size != -1 {
		ossOpts = append(ossOpts, oss.ContentLength(opt.Size))
	}
	if opt.ContentType != "" {
		ossOpts = append(ossOpts, oss.ContentType(opt.ContentType))
	}

	if err := c.bucket.PutObject(name, reader, ossOpts...); err != nil {
		return errors.Wrapf(err, "oss put object err")
	}

	return nil
}

func (c *ossClient) RemoveObject(ctx context.Context, name string) error {
	if err := c.bucket.DeleteObject(name); err != nil {
		return errors.Wrapf(err, "oss delete object err")
	}
	return nil
}

func (c *ossClient) SignedURL(ctx context.Context, name string, method string, expires time.Duration) (*url.URL, error) {
	raw, err := c.bucket.SignURL(name, oss.HTTPMethod(method), int64(expires.Seconds()), nil)
	if err != nil {
		return nil, errors.Wrapf(err, "oss signed url err")
	}

	u, err := url.Parse(raw)
	if err != nil {
		return nil, errors.Wrapf(err, "oss parse raw url err")
	}

	return u, nil
}

var _ Client = (*ossClient)(nil)

func NewOSSClient(endpoint string, accessKey string, secretKey string, bucketName string) (*ossClient, error) {
	if bucketName == "" {
		return nil, errors.New("oss bucket is empty")
	}

	// endpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// accessKey和secretKey填写RAM用户的访问密钥（AccessKey ID和AccessKey Secret）。
	client, err := oss.New(endpoint, accessKey, secretKey)
	if err != nil {
		return nil, errors.Wrapf(err, "oss new client err")
	}

	exist, err := client.IsBucketExist(bucketName)
	if err != nil {
		return nil, errors.Wrapf(err, "oss bucket exists err")
	}
	if !exist {
		err = client.CreateBucket(bucketName)
		if err != nil {
			return nil, errors.Wrapf(err, "oss create bucket err")
		}
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, errors.Wrapf(err, "oss get bucket err")
	}

	return &ossClient{
		kernel: client,
		bucket: bucket,
	}, nil
}

func (c *ossClient) Kernel() *oss.Client {
	return c.kernel
}
