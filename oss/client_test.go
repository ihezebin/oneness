package oss

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestMinioClient(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient("minio://CscaPrbfoRO0EapeM56m:RNxYKk9XtdIAyFnCHNLnyYei2tmX4L59ACeDe0Ap@127.0.0.1:9000/test")
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.Open("./test.txt")
	if err != nil {
		t.Fatal(err)
	}

	key := filepath.Base(file.Name())

	err = client.PutObject(ctx, key, file)
	if err != nil {
		t.Fatal(err)
	}

	info, err := client.StatObject(ctx, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", info)

	u, err := client.SignedURL(ctx, key, http.MethodGet, time.Minute)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", u)

	err = client.RemoveObject(ctx, key)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOSSClient(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient("oss://LTAI5tRrwfxuS3TYvawJixRu:jXnXHt22oDtwIVhQSrMXCyiRamG92q@oss-cn-chengdu.aliyuncs.com/oneness-test")
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.Open("./test.txt")
	if err != nil {
		t.Fatal(err)
	}

	key := filepath.Base(file.Name())

	err = client.PutObject(ctx, key, file)
	if err != nil {
		t.Fatal(err)
	}

	info, err := client.StatObject(ctx, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", info)

	err = client.RemoveObject(ctx, key)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCosClient(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient("cos://AKIDpaXWddWVYTB5whEb0LlCVklYBjCXu8B6:t4B6uQpGrQ9lFcmppiVFpZOPLAoV4seH@cos.ap-chengdu.myqcloud.com/test-1258606727")
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.Open("./test.txt")
	if err != nil {
		t.Fatal(err)
	}

	key := filepath.Base(file.Name())

	err = client.PutObject(ctx, key, file)
	if err != nil {
		t.Fatal(err)
	}

	info, err := client.StatObject(ctx, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", info)

	err = client.RemoveObject(ctx, key)
	if err != nil {
		t.Fatal(err)
	}

}

func TestKodoClient(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient("kodo://alWFYoSm_BnGy6Wt-XAVFNiSrp4-7ARRTVTPV8pr:CNp5osmW0OKdTkwts0MKGN0qHFTf35LIe6Ys2TFo@z2/oneness-test")
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.Open("./test.txt")
	if err != nil {
		t.Fatal(err)
	}

	key := filepath.Base(file.Name())

	err = client.PutObject(ctx, key, file)
	if err != nil {
		t.Fatal(err)
	}

	info, err := client.StatObject(ctx, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", info)

	err = client.RemoveObject(ctx, key)
	if err != nil {
		t.Fatal(err)
	}

}
