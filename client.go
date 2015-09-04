package whitepages

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	libraryVersion = "0.1"
	defaultBaseURL = "https://proapi.whitepages.com/"
	version        = "2.1"
	defaultTimeout = 10 * time.Second
)

type Client struct {
	ApiKey    string
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
	c.ApiKey = key
	baseURL := defaultBaseURL + version + "/"
	c.BaseURL = baseURL
	return c
}

func (c *Client) Phone(params map[string]string, timeout time.Duration) (PhoneResponse, error) {
	p := PhoneResponse{}
	response, err := c.request("phone.json", timeout, params)
	if err != nil {
		return p, err
	}
	if err = json.Unmarshal(response, &p); err != nil {
		return p, err
	}
	return p, err

}

func (c *Client) Address(params map[string]string, timeout time.Duration) (AddressResponse, error) {
	p := AddressResponse{}
	response, err := c.request("location.json", timeout, params)
	if err != nil {
		return p, err
	}
	if err = json.Unmarshal(response, &p); err != nil {
		return p, err
	}
	return p, err

}

func (c *Client) request(method string, timeout time.Duration, params map[string]string) ([]byte, error) {
	req, _ := url.Parse(c.BaseURL + method)
	p := url.Values{}

	for k, v := range params {
		p.Add(k, v)
	}
	p.Add("api_key", c.ApiKey)
	req.RawQuery = p.Encode()

	// maybe move http client up into struct
	client := http.Client{
		Timeout: timeout,
	}
	response, err := client.Get(req.String())
	log.Println(req.String())

	if err != nil {
		return []byte{}, err
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return []byte{}, err
		}
		e := ErrorResponse{}
		if err = json.Unmarshal(contents, &e); err != nil {
			return nil, err
		}
		if len(e.Error.Message) > 0 {
			return nil, errors.New(e.Error.Message)
		}
		return contents, err
	}
}
