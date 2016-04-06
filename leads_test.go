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

// "request": {
//   "address.postal_code": "59004",
//   "address.state": "MT",
//   "address.street_line_1": "302 Gorham Ave",
//   "address_city": "Ashland",
//   "api_key": "3a83ab40280b6926ee62bc50b4de64c4",
//   "email_address": "medjalloh1@yahoo.com",
//   "ip_address": "108.194.128.165",
//   "name": "Drama Number",
//   "phone": "6464806649"
// }
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

// name=Drama+Number&phone=6464806649&email_address=medjalloh1@yahoo.com&address.street_line_1=302+Gorham+Ave&
// address_city=Ashland&address.postal_code=59004&address.state=MT&ip_address=108.194.128.165
// https://proapi.whitepages.com/3.1/lead_verify_append.json?api_key=3bc9ee9e70576cc0aff8c0fe38a47b5d&firstname=Chester&lastname=Stevens&email_address=dabearsck2013@comcast.net&phone=3095319550&address.street_line_1=706+N.+Lee+St.&address.postal_code=61701&address_city=Bloomington&address.state=IL

// https://proapi.whitepages.com/3.1/lead_verify.json?api_key=3a83ab40280b6926ee62bc50b4de64c4&firstname=Chester&lastname=Stevens&email_address=dabearsck2013@comcast.net&phone=3095319550&address.street_line_1=706+N.+Lee+St.&address.postal_code=61701&address_city=Bloomington&address.state=IL
