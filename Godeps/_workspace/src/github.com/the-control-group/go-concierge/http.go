package concierge

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HTTPDoer does requests over HTTP
type HTTPDoer struct {
	request       *http.Request
	requestBytes  []byte
	response      *http.Response
	responseBytes []byte
}

// NewHTTPDoer creates a new HTTPDoer with the provided request
func NewHTTPDoer(request *http.Request) *HTTPDoer {
	doer := new(HTTPDoer)
	doer.request = request

	buf := new(bytes.Buffer)
	err := request.Write(buf)
	if err == nil {
		doer.requestBytes = buf.Bytes()
	}

	return doer
}

// GetRequest returns the original request object
func (doer *HTTPDoer) GetRequest() interface{} {
	return doer.request
}

// GetRequestBytes returns the request bytes
func (doer *HTTPDoer) GetRequestBytes() []byte {
	return doer.requestBytes
}

// GetResponse returns the response object or nil if doer hasn't been done yet
func (doer *HTTPDoer) GetResponse() interface{} {
	return doer.response
}

// GetResponseBytes returns the response bytes
func (doer *HTTPDoer) GetResponseBytes() []byte {
	return doer.responseBytes
}

// ResponseError returns an error if the HTTP status was not OK
func (doer *HTTPDoer) ResponseError() error {
	if doer.response == nil {
		return fmt.Errorf("response is nil")
	}
	if doer.response.StatusCode != http.StatusOK {
		return fmt.Errorf("expected status OK, but got '%s'", doer.response.Status)
	}

	return nil
}

// AddStats adds HTTP specific stats
func (doer *HTTPDoer) AddStats(stats map[string]interface{}) {
	if doer.response != nil {
		stats["status_code"] = doer.response.StatusCode
	}

	return
}

// Do creates an http.Client and runs Do
func (doer *HTTPDoer) Do(timeout time.Duration) error {
	client := new(http.Client)
	client.Timeout = timeout

	response, err := client.Do(doer.request)
	doer.response = response
	if err != nil {
		return err
	}
	defer doer.response.Body.Close()

	doer.responseBytes, err = ioutil.ReadAll(doer.response.Body)
	if err != nil {
		return err
	}

	return nil
}
