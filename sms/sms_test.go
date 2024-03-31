package smsc

import (
	"testing"

	"github.com/ihezebin/oneness/sms/aliyun"
	"github.com/ihezebin/oneness/sms/tencent"
	"github.com/ihezebin/sdk/utils/common"
)

func TestTencentSms(t *testing.T) {
	client, err := tencent.NewClientWithConfig(tencent.Config{
		SecretId:  "SecretId",
		SecretKey: "SecretKey",
		Region:    "ap-guangzhou",
	})
	if err != nil {
		t.Fatal(err)
	}
	msg := tencent.NewMessage().WithAppId("1400578890").WithSignName("hezebin").
		WithTemplate("11477481", 123321, 10)
	faileds, err := client.SendSms(msg, "+8613518468111")
	if err != nil {
		t.Error(faileds)
		t.Fatal(err)
	}
	t.Log("send sms succeed")
}

func TestAliyunSms(t *testing.T) {
	client, err := aliyun.NewClientWithConfig(aliyun.Config{
		AccessKeyId:     "",
		AccessKeySecret: "",
	})
	if err != nil {
		t.Fatal(err)
	}
	msg := aliyun.NewMessage().WithSignName("sign").WithTemplate("code", common.Json{})
	err = client.SendSms(msg, "13518468111")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("send sms succeed")
}
