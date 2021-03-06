package whitepages

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/the-control-group/go-concierge"
)

var client *Client
var key, lvKey, lvaKey string

func init() {
	key = os.Getenv("WP_API_KEY")
	lvKey = os.Getenv("WP_LV_API_KEY")
	lvaKey = os.Getenv("WP_LVA_API_KEY")
	client = NewClient(key)
}

func TestKey(t *testing.T) {
	if client.APIKey == "" {
		t.Error("api key not set, you must set the environment variable WP_API_KEY before running")
		t.FailNow()
	}
}

func TestMarshalSampleResponse(t *testing.T) {
	Convey("Given a sample response", t, func() {
		file, err := ioutil.ReadFile("./fixtures/sample_phone_response.json")
		if err != nil {
			t.Errorf("%s", err)
		}

		var resp PhoneResponse
		Convey("When unmarshaling it should not cause any errors", func() {
			err = json.Unmarshal(file, &resp)
			So(err, ShouldBeNil)
		})
	})
}

func TestPhone(t *testing.T) {
	Convey("Test phone lite", t, func() {
		timeout := time.Duration(20 * time.Second)
		params := make(map[string]string)
		params["phone_number"] = "2069735100"
		params["response_type"] = "lite"
		options := concierge.Options{Timeout: timeout}

		_, _, err := client.Phone(params, options)
		So(err, ShouldBeNil)

	})

}

func TestLandlinePhone(t *testing.T) {
	Convey("Test phone full with landline", t, func() {
		timeout := time.Duration(20 * time.Second)
		params := make(map[string]string)
		params["phone_number"] = "5169389674"
		options := concierge.Options{Timeout: timeout}

		resp, _, err := client.Phone(params, options)
		So(err, ShouldBeNil)
		So(resp.Results, ShouldNotBeEmpty)
	})
}
