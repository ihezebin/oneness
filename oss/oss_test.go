package oss

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/ihezebin/oneness/oss/qiniu"
	"github.com/ihezebin/oneness/oss/tencent"
	"github.com/ihezebin/oneness/oss/ucloud"
)

func TestTencent(t *testing.T) {
	ctx := context.Background()
	client := tencent.NewClientWithConfig(tencent.Config{
		SecretID:  "SecretID",
		SecretKey: "SecretKey",
		BucketURL: "http://images.hezebin.com",
	})
	url, err := client.Upload(ctx, strings.NewReader("test file3"), "test3.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(url)
	err = client.Delete(ctx, "test.txt")
	if err != nil {
		t.Fatal(err)
	}
	faileds, err := client.DeleteMulti(ctx, "test1.txt", "test2.txt", "test3.txt")
	if err != nil {
		t.Fatal(err)
	}
	if len(faileds) > 0 {
		t.Error(faileds)
	}
	t.Log("upload succeed")
}

func TestQiniu(t *testing.T) {
	client := qiniu.NewClientWithConfig(qiniu.Config{
		Zone:      qiniu.ZoneHuanan,
		AccessKey: "AccessKey",
		SecretKey: "SecretKey",
		Bucket:    "c4lms",
		Domain:    "http://image-c4lms-qiniu.whereabouts.icu",
	})
	file, err := os.Open("C:\\Users\\Korbin\\Pictures\\hzb.jpg")
	if err != nil {
		t.Log(err)
		return
	}
	url, err := client.Upload(file, "Korbin.jpg")
	if err != nil {
		t.Log(err, url)
		return
	}
	fmt.Println(url)
}

func TestUCloud(t *testing.T) {
	client := ucloud.NewClientWithConfig(ucloud.Config{
		PublicKey:  "PublicKey",
		PrivateKey: "PrivateKey",
		FileHost:   "cn-bj.ufileos.com",
		BucketName: "c4lms",
	})
	file, err := os.Open("C:\\Users\\Korbin\\Pictures\\hzb.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	url, err := client.Upload(file, "Korbin.jpg")
	if err != nil {
		t.Log(err, url)
		return
	}
	fmt.Println(url)
}
