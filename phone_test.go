package whitepages

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
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

func TestMarshalSampleResponse(t *testing.T) {

	Convey("Given a sample response", t, func() {
		file, err := ioutil.ReadFile("./sample_phone_response.json")
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

		_, err := client.Phone(params, timeout)
		So(err, ShouldBeNil)

	})

}

func TestLandlinePhone(t *testing.T) {
	Convey("Test phone full with landline", t, func() {
		timeout := time.Duration(20 * time.Second)
		params := make(map[string]string)
		params["phone_number"] = "5169389674"

		resp, err := client.Phone(params, timeout)
		So(err, ShouldBeNil)
		So(resp.Results, ShouldNotBeEmpty)
	})
}
