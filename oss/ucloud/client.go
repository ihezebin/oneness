package ucloud

import (
	"context"
	"io"

	ufsdk "github.com/ufilesdk-dev/ufile-gosdk"
)

type Client struct {
	kernel *ufsdk.Config
	config Config
}

func NewClient(config Config) *Client {
	return &Client{
		kernel: &ufsdk.Config{
			PublicKey:       config.PublicKey,
			PrivateKey:      config.PrivateKey,
			BucketName:      config.BucketName,
			BucketHost:      config.BucketHost,
			FileHost:        config.FileHost,
			VerifyUploadMD5: config.VerifyUploadMD5,
		},
		config: config,
	}
}

func (client *Client) Kernel() *ufsdk.Config {
	return client.kernel
}

func (client *Client) Upload(ctx context.Context, file io.Reader, key string) (string, error) {
	req, err := ufsdk.NewFileRequest(client.kernel, nil)
	if err != nil {
		return "", err
	}
	err = req.IOPut(file, key, "")
	if err != nil {
		return string(req.DumpResponse(true)), err
	}
	return req.GetPublicURL(key), err
}

func (client *Client) Delete(ctx context.Context, key string) error {
	req, err := ufsdk.NewFileRequest(client.kernel, nil)
	if err != nil {
		return err
	}
	err = req.DeleteFile(key)
	if err != nil {
		return err
	}
	return nil
}
