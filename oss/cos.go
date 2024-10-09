package oss

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/tencentyun/cos-go-sdk-v5"
)

// https://cloud.tencent.com/document/product/436/31215
type cosClient struct {
	kernel    *cos.Client
	bucket    string
	accessKey string
	secretKey string
}

func (c *cosClient) StatObject(ctx context.Context, name string) (*ObjectInfo, error) {
	resp, err := c.kernel.Object.Head(ctx, name, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "cos head object err")
	}
	defer resp.Body.Close()

	info := &ObjectInfo{
		ContentType: resp.Header.Get("Content-Type"),
		ETag:        resp.Header.Get("Etag"),
	}
	if v := resp.Header.Get("Content-Length"); v != "" {
		info.Size, _ = strconv.ParseInt(v, 10, 64)
	}
	if v := resp.Header.Get("Last-Modified"); v != "" {
		info.LastModified, _ = time.Parse(time.RFC1123, v)
	}

	return info, nil
}

func (c *cosClient) GetObjects(ctx context.Context, prefix string, opts ...GetObjectsOption) ([]Object, error) {
	prefix += "/"
	prefix = strings.ReplaceAll(prefix, "//", "/")

	opt := newGetObjectsOptions(opts...)

	objs := make([]Object, 0)

	isTruncated := true
	marker := ""
	for isTruncated {
		v, _, err := c.kernel.Bucket.Get(ctx, &cos.BucketGetOptions{
			Prefix:    prefix,
			MaxKeys:   opt.Limit,
			Marker:    marker,
			Delimiter: "/", // 固定按目录查找
		})
		if err != nil {
			return nil, errors.Wrapf(err, "cos get objects err, marker: %v, isTruncated: %v", marker, isTruncated)
		}
		for _, content := range v.Contents {
			key := content.Key
			object, err := c.GetObject(ctx, key)
			if err != nil {
				return nil, errors.Wrapf(err, "cos get object err")
			}

			objs = append(objs, Object{
				Key:  key,
				Data: object,
			})
		}
		isTruncated = v.IsTruncated // 是否还有数据
		marker = v.NextMarker       // 设置下次请求的起始 key
	}

	return objs, nil
}

func (c *cosClient) GetObject(ctx context.Context, name string) (io.ReadCloser, error) {
	resp, err := c.kernel.Object.Get(ctx, name, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "cos get object err")
	}

	return resp.Body, nil
}

func (c *cosClient) PutObject(ctx context.Context, name string, reader io.Reader, opts ...PutOption) error {
	opt := newPutOptions(opts...)

	cosOpts := &cos.ObjectPutOptions{
		ACLHeaderOptions:       &cos.ACLHeaderOptions{},
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{},
	}

	if opt.Size != -1 {
		cosOpts.ContentLength = opt.Size
	}

	if opt.ContentType != "" {
		cosOpts.ContentType = opt.ContentType
	}

	if _, err := c.kernel.Object.Put(ctx, name, reader, cosOpts); err != nil {
		return errors.Wrapf(err, "cos put object err")
	}

	return nil
}

func (c *cosClient) RemoveObject(ctx context.Context, name string) error {
	_, err := c.kernel.Object.Delete(ctx, name)
	if err != nil {
		return errors.Wrapf(err, "cos remove object err")
	}
	return nil
}

func (c *cosClient) SignedURL(ctx context.Context, name string, method string, expires time.Duration) (*url.URL, error) {
	u, err := c.kernel.Object.GetPresignedURL(ctx, method, name, c.accessKey, c.secretKey, expires, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "cos signed url err")
	}

	return u, nil
}

var _ Client = (*cosClient)(nil)

func NewCOSClient(endpoint string, accessKey string, secretKey string, bucket string) (*cosClient, error) {
	if bucket == "" {
		return nil, errors.New("minio bucket is empty")
	}

	regions := regexp.MustCompile(`cos\.([^.]+)\.myqcloud\.com`).FindStringSubmatch(endpoint)
	if len(regions) != 2 {
		return nil, errors.New("cos endpoint is invalid, example: cos.ap-shanghai.myqcloud.com")
	}

	bucketUrl, err := cos.NewBucketURL(bucket, regions[1], true)
	if err != nil {
		return nil, errors.Wrapf(err, "new cos bucket url err")
	}

	client := cos.NewClient(&cos.BaseURL{BucketURL: bucketUrl}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  accessKey,
			SecretKey: secretKey,
		},
	})

	ctx := context.Background()
	exist, err := client.Bucket.IsExist(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "cos bucket is exist err")
	}
	if !exist {
		_, err = client.Bucket.Put(ctx, &cos.BucketPutOptions{})
		if err != nil {
			return nil, errors.Wrapf(err, "make cos bucket err")
		}
	}

	return &cosClient{
		kernel:    client,
		bucket:    bucket,
		accessKey: accessKey,
		secretKey: secretKey,
	}, nil
}

func (c *cosClient) Kernel() *cos.Client {
	return c.kernel
}
