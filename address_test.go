package whitepages

import (
	"testing"
	"time"
)

func TestAddress(t *testing.T) {
	timeout := time.Duration(20 * time.Second)
	params := make(map[string]string)
	params["street_line_1"] = "1301 5th Avenue"
	params["city"] = "Seatle"
	params["state"] = "WA"

	err, _ := client.Address(params, timeout)
	if err != nil {
		t.Error(err)
	}
}
