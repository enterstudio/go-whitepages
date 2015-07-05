package whitepages

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	libraryVersion = "0.1"
	defaultBaseURL = "https://proapi.whitepages.com/"
	v2Version      = "2.1"
	defaultTimeout = 10 * time.Second
)

type V2Client struct {
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

func NewV2Client(key string) *V2Client {
	c := &V2Client{}
	c.ApiKey = key
	baseURL := defaultBaseURL + v2Version + "/"
	c.BaseURL = baseURL
	return c
}

func (c *V2Client) request(method string, timeout time.Duration, params map[string]string) (error, []byte) {
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

	if err != nil {
		return err, []byte{}
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err, []byte{}
		}
		e := ErrorResponse{}
		if err = json.Unmarshal(contents, &e); err != nil {
			return err, nil
		}
		if len(e.Error.Message) > 0 {
			return errors.New(e.Error.Message), nil
		}
		return nil, contents
	}
}
