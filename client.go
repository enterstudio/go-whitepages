package whitepages

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/the-control-group/go-concierge"
)

const (
	libraryVersion = "0.1"
	defaultBaseURL = "https://proapi.whitepages.com/"
	version        = "2.1"
	defaultTimeout = 10 * time.Second
)

// Client is the whitepages client
type Client struct {
	APIKey    string
	Version   string
	BaseURL   string
	UserAgent string
	Timeout   time.Duration
}

// ErrorResponse is the differently structured ErrorResponse. This can probably be set in the actual response so either one that returns will be populated.
type ErrorResponse struct {
	Error struct {
		Message string `json:"message"`
		Name    string `json:"name"`
	} `json:"error"`
}

// NewClient is a client constructor method
func NewClient(key string) *Client {
	c := &Client{}
	c.APIKey = key
	baseURL := defaultBaseURL + version + "/"
	c.BaseURL = baseURL
	return c
}

// Lead verify
// Request variables:api_key, name, first_name, last_name, phone, email_address, ip_address,address.city, address.state_code,address.county_code
// 1. Create method
// 2. Create structs
// 3. Tests
func (c *Client) LeadVerify(params map[string]string, opts concierge.Options) (l LeadVerifyResponse, cached bool, err error) {
	response, cached, err := c.request("lead_verify.json", params, opts)
	if err != nil {
		return
	}
	if err = json.Unmarshal(response, &l); err != nil {
		return
	}
	return
}

// Phone is a reverse phone search
func (c *Client) Phone(params map[string]string, opts concierge.Options) (PhoneResponse, bool, error) {
	p := PhoneResponse{}
	response, cached, err := c.request("phone.json", params, opts)
	if err != nil {
		return p, cached, err
	}
	if err = json.Unmarshal(response, &p); err != nil {
		return p, cached, err
	}
	return p, cached, err
}

// Address is a reverse location search
func (c *Client) Address(params map[string]string, opts concierge.Options) (AddressResponse, bool, error) {
	p := AddressResponse{}
	response, cached, err := c.request("location.json", params, opts)
	if err != nil {
		return p, cached, err
	}
	if err = json.Unmarshal(response, &p); err != nil {
		return p, cached, err
	}
	return p, cached, err
}

// probably no need for this method. Do it all in the above with some common code refactored into functions
func (c *Client) request(method string, params map[string]string, opts concierge.Options) (res []byte, cached bool, err error) {
	// Build URL
	req, _ := url.Parse(c.BaseURL + method)
	p := url.Values{}

	for k, v := range params {
		p.Add(k, v)
	}
	p.Add("api_key", c.APIKey)
	req.RawQuery = p.Encode()

	// Build request
	request, err := http.NewRequest("GET", req.String(), nil)
	if err != nil {
		return
	}

	doer := concierge.NewHTTPDoer(request)
	res, cached, err = concierge.Request(method, doer, opts)
	if err != nil {
		return
	}

	return
}

// 	if err != nil {
// 		return []byte{}, err
// 	} else {
// 		defer response.Body.Close()
// 		contents, err := ioutil.ReadAll(response.Body)
// 		if err != nil {
// 			return []byte{}, err
// 		}
// 		e := ErrorResponse{}
// 		if err = json.Unmarshal(contents, &e); err != nil {
// 			return nil, err
// 		}
// 		if len(e.Error.Message) > 0 {
// 			return nil, errors.New(e.Error.Message)
// 		}
// 		return contents, err
// 	}
// }
