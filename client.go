package whitepages

import (
	"encoding/json"
	"errors"
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

type Client struct {
	APIKey    string
	Version   string
	BaseURL   string
	UserAgent string
	Timeout   time.Duration
}

type ErrorResponse struct {
	Error struct {
		Message string `json:"message"`
		Name    string `json:"name"`
	} `json:"error"`
}

func NewClient(key string) *Client {
	c := &Client{}
	c.APIKey = key
	baseURL := defaultBaseURL + version + "/"
	c.BaseURL = baseURL
	return c
}

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

	// defer response.Body.Close()

	// res, err = ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	return nil, err
	// }

	e := ErrorResponse{}
	err = json.Unmarshal(res, &e)
	if err != nil {
		return
	}

	if len(e.Error.Message) > 0 {
		err = errors.New(e.Error.Message)
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
