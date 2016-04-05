package whitepages

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMarshalSampleLeadVerifyResponse(t *testing.T) {
	Convey("Given a sample response", t, func() {
		file, err := ioutil.ReadFile("./fixtures/sample_lead_verify_response.json")
		if err != nil {
			t.Errorf("%s", err)
		}

		var resp LeadVerifyResponse
		Convey("When unmarshaling it should not cause any errors", func() {
			err = json.Unmarshal(file, &resp)
			So(err, ShouldBeNil)
		})
	})
}

// func TestPhone(t *testing.T) {
// 	Convey("Test phone lite", t, func() {
// 		timeout := time.Duration(20 * time.Second)
// 		params := make(map[string]string)
// 		params["phone_number"] = "2069735100"
// 		params["response_type"] = "lite"
// 		options := concierge.Options{Timeout: timeout}

// 		_, _, err := client.Phone(params, options)
// 		So(err, ShouldBeNil)

// 	})
// }
