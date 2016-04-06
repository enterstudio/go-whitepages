package whitepages

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/the-control-group/go-concierge"
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

func TestMarshalSampleLeadVerifyAppendResponse(t *testing.T) {
	Convey("Given a sample response", t, func() {
		file, err := ioutil.ReadFile("./fixtures/sample_lead_verify_append_response.json")
		if err != nil {
			t.Errorf("%s", err)
		}

		var resp LeadVerifyResponse
		Convey("When unmarshaling it should not cause any errors", func() {
			err = json.Unmarshal(file, &resp)
			So(err, ShouldBeNil)
			So(resp.PhoneChecks.IsDoNotCallRegistered, ShouldNotBeNil)
		})
	})
}

func TestLeadVerify(t *testing.T) {
	Convey("Test lead verify", t, func() {
		timeout := time.Duration(20 * time.Second)
		params := make(map[string]string)
		params["phone"] = "6464806649"
		params["api_key"] = lvKey
		options := concierge.Options{Timeout: timeout}
		response, _, err := client.LeadVerify(params, options)
		So(response.AddressChecks, ShouldBeNil)
		So(response.EmailAddressChecks, ShouldBeNil)
		So(response.IPAddressChecks, ShouldBeNil)
		So(response.NameChecks, ShouldBeNil)
		So(response.PhoneChecks, ShouldNotBeNil)
		So(response.PhoneChecks.Error, ShouldBeNil)
		So(response.PhoneChecks.IsConnected, ShouldBeNil)
		So(err, ShouldBeNil)
	})
}

func TestLeadVerifyAppend(t *testing.T) {
	Convey("Test lead verify", t, func() {
		timeout := time.Duration(20 * time.Second)
		params := make(map[string]string)
		params["phone"] = "6464806649"
		params["firstname"] = "Drama"
		params["lastname"] = "Number"
		params["api_key"] = lvaKey
		options := concierge.Options{Timeout: timeout}
		response, _, err := client.LeadVerifyAppend(params, options)
		t.Logf("%+v", response)
		So(err, ShouldBeNil)
		So(response.AddressChecks, ShouldBeNil)
		So(response.EmailAddressChecks, ShouldBeNil)
		So(response.IPAddressChecks, ShouldBeNil)
		So(response.NameChecks, ShouldNotBeNil)
		So(response.PhoneChecks, ShouldNotBeNil)
		So(response.PhoneChecks.Error, ShouldBeNil)
		So(response.PhoneChecks.IsConnected, ShouldBeNil)
		So(response.PhoneChecks.IsPrepaid, ShouldBeNil)
		So(response.PhoneChecks.IsDoNotCallRegistered, ShouldNotBeNil)
		So(*response.PhoneChecks.IsDoNotCallRegistered, ShouldBeTrue)
	})
}
