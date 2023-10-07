package resty_test

import (
	"testing"

	"github.com/go-resty/resty/v2"

	"github.com/nevercatchmeow/penumbra/core/tools/log"
)

func TestResty(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		t.Fatal(err)
	}
	log.Info("resp", log.ByteString("bytes", resp.Body()))
}
