package whitepages

import (
	"os"
	"testing"
	"time"
)

var client *Client

func init() {
	key := os.Getenv("WP_API_KEY")
	client = NewClient(key)
}

func TestKey(t *testing.T) {
	if client.ApiKey == "" {
		t.Error("api key not set, you must set the environment variable WP_API_KEY before running")
		t.FailNow()
	}
}

func TestPhone(t *testing.T) {
	timeout := time.Duration(20 * time.Second)
	params := make(map[string]string)
	params["phone_number"] = "2069735100"
	params["response_type"] = "lite"

	err, _ := client.Phone(params, timeout)
	if err != nil {
		t.Error(err)
	}
}
