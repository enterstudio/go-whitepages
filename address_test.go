package whitepages

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/the-control-group/go-concierge"
)

func TestAddress(t *testing.T) {
	Convey("Test the address method", t, func() {
		timeout := time.Duration(20 * time.Second)
		options := concierge.Options{
			Timeout: timeout,
		}

		params := make(map[string]string)
		params["street_line_1"] = "1301 5th Avenue"
		params["city"] = "Seatle"
		params["state"] = "WA"

		_, _, err := client.Address(params, options)
		So(err, ShouldNotBeNil)
		errString := "concierge.Request(): response not OK for 'location.json': expected status OK, but got '403 Forbidden"
		So(err.Error(), ShouldContainSubstring, errString)

		// So(addr.Results, ShouldNotBeEmpty)
		// for _, v := range addr.Results {
		// 	log.Printf("%+v", v)
		// }
	})

}

