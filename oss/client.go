package oss

import (
	"context"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	SchemeOSS   = "oss" // 阿里云
	SchemeCOS   = "cos" // 腾讯云
	SchemeMinio = "minio"
	SchemeKodo  = "kodo" // 七牛云
	SchemeUS3   = "us3"  // ucloud
)

type ObjectInfo struct {
	Key          string
	ETag         string
	Size         int64
	ContentType  string
	LastModified time.Time
	Expires      time.Time
}

type PutOptions struct {
	Size        int64
	ContentType string
}

type PutOption func(*PutOptions)

func WithSize(size int64) PutOption {
	return func(opt *PutOptions) {
		opt.Size = size
	}
}

func WithContentType(contentType string) PutOption {
	return func(opt *PutOptions) {
		opt.ContentType = contentType
	}
}

func newPutOptions(opts ...PutOption) *PutOptions {
	opt := &PutOptions{
		Size: -1,
	}
	for _, optFunc := range opts {
		optFunc(opt)
	}
	return opt
}

type Client interface {
	StatObject(ctx context.Context, name string) (*ObjectInfo, error)

	GetObject(ctx context.Context, name string) (io.ReadCloser, error)

	PutObject(ctx context.Context, name string, reader io.Reader, opts ...PutOption) error

	RemoveObject(ctx context.Context, name string) error

	SignedURL(ctx context.Context, name string, method string, expires time.Duration) (*url.URL, error)
}

func NewClient(dsn string) (Client, error) {
	u, err := url.Parse(dsn)
	if err != nil {
		return nil, errors.Wrapf(err, "parse dsn err")
	}

	if u.User == nil {
		return nil, errors.New("missing access key or secret")
	}

	bucket := strings.Trim(u.Path, "/")
	if bucket == "" {
		return nil, errors.New("missing bucket")
	}

	accessKey := u.User.Username()
	secretKey, _ := u.User.Password()

	switch strings.ToLower(u.Scheme) {
	case SchemeMinio:
		return NewMinioClient(u.Host, accessKey, secretKey, bucket)
	case SchemeOSS:
		return NewOSSClient(u.Host, accessKey, secretKey, bucket)
	case SchemeCOS:
		return NewCOSClient(u.Host, accessKey, secretKey, bucket)
	//case SchemeUS3:
	//case SchemeKodo:
	//return NewKodoClient(u.Host, accessKey, secretKey, bucket)
	default:
		return nil, errors.Errorf("unsupported schema: %s", u.Scheme)
	}
}
