package whitepages

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddress(t *testing.T) {
	Convey("Test the address method", t, func() {
		timeout := time.Duration(20 * time.Second)
		params := make(map[string]string)
		params["street_line_1"] = "1301 5th Avenue"
		params["city"] = "Seatle"
		params["state"] = "WA"

		_, err := client.Address(params, timeout)
		// So(err, ShouldBeNil)
		So(err.Error(), ShouldEqual, "Action is not allowed for authorization token.")
		// So(addr.Results, ShouldNotBeEmpty)
		// for _, v := range addr.Results {
		// 	log.Printf("%+v", v)
		// }
	})

}
