package httpclient

import (
	"context"
	"testing"
)

var ctx = context.Background()

func TestClient(t *testing.T) {
	response, err := NewRequest(ctx).Get("https://baidu.com")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", response.Body())
}
