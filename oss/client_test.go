package oss

import (
	"bytes"
	"context"
	"io"
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

	file, err := os.Open("./664b73a08543301506115cf8.jpg")
	if err != nil {
		t.Fatal(err)
	}

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		t.Fatal(err)
	}

	contentType := http.DetectContentType(buffer)
	_, err = file.Seek(0, 0)
	if err != nil {
		t.Fatal(err)
	}

	stat, err := file.Stat()
	if err != nil {
		t.Fatal(err)
	}

	key := filepath.Base(file.Name())
	err = client.PutObject(ctx, key, file, WithContentType(contentType), WithSize(stat.Size()))
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
	client, err := NewClient("oss://xxx:xxx@oss-cn-chengdu.aliyuncs.com/oneness-test")
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

func TestCosClientGetObject(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient("cos://xxx:xxx@cos.ap-chengdu.myqcloud.com/test-1258606727")
	if err != nil {
		t.Fatal(err)
	}

	err = client.PutObject(ctx, "test_data", bytes.NewReader([]byte("test_data")))
	if err != nil {
		t.Fatal(err)
	}

	object, err := client.GetObject(ctx, "test_data")
	if err != nil {
		t.Fatal(err)
	}

	data, err := io.ReadAll(object)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", data)
}

func TestCosClientGetObjects(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient("cos://xxx:xxx@cos.ap-chengdu.myqcloud.com/test-1258606727")
	if err != nil {
		t.Fatal(err)
	}

	objects, err := client.GetObjects(ctx, "backup_test/blog-minio/preview/")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%d", len(objects))

	for i, object := range objects {
		data, err := io.ReadAll(object.Data)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%s", data)
		t.Logf("%d", i)
	}
}

func TestCosClient(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient("cos://xxx:xxxx@cos.ap-chengdu.myqcloud.com/test-1258606727")
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
	client, err := NewClient("kodo://xx-xxx-7ARRTVTPV8pr:xxx@z2/oneness-test")
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
