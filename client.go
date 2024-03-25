package gowrt

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	url         url.URL
	httpClient  *http.Client
	ubusSession UbusSession
}

type Option func(*Client)

func WithTimeout(t time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = t
	}
}

func WithHttpTransport(t time.Duration) Option {
	return func(c *Client) {
		c.url.Scheme = "http"
	}
}

func WithInsecureTls() Option {
	return func(c *Client) {
		c.httpClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}
}

func New(host string, options ...Option) *Client {
	client := &Client{
		url: url.URL{
			Scheme: "https",
			Host:   host,
			Path:   "/ubus",
		},
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		ubusSession: UbusSession{
			UbusRpcSession: "00000000000000000000000000000000",
		},
	}
	for _, opt := range options {
		opt(client)
	}
	return client
}

func (c *Client) ApiCall(call rpcCall) ([]byte, error) {
	var response rpcResponse
	// ensure call is valid
	err := call.validate()
	if err != nil {
		return nil, fmt.Errorf("invalid call: %s", err)
	}
	jsonBody, err := json.Marshal(call.toApiPayload(c.ubusSession))
	if err != nil {
		return nil, fmt.Errorf("marshal api request: %s", err)
	}
	req, err := http.NewRequest(http.MethodPost, c.url.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("create api request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "gowrt v0.0.1")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send api request: %s", err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("decode api response: %s", err)
	}
	// FIXME: response errors
	// if response.Error != nil {
	// 	return nil, fmt.Errorf("response error %d: %s", response.Error["code"])
	// }

	// The API returns a response like below:
	// {
	// 	"jsonrpc": "2.0",
	// 	"id": 2,
	// 	"result": [
	// 	  0,
	// 	  {
	// 		"values":  ... data we care about
	//
	// result is a two-item list where the first value is an rpc/ubus status code
	// that should be checked, and the second entry is the actual data the
	// caller cares about.
	for i, r := range response.Result {
		switch i {
		case 0:
			if r != float64(0) {
				return nil, fmt.Errorf("call failed: unexpected response: %d", r)
			}
		case 1:
			// Put the response back into json for the caller to
			// unmarshal into the proper struct type
			responseData, err := json.Marshal(r)
			if err != nil {
				return nil, fmt.Errorf("remarshal api response: %s", err)
			}
			return responseData, nil
		default:
			return nil, fmt.Errorf("unexpected additional response: %s", r)
		}
	}
	// should never reach here
	return nil, fmt.Errorf("unknown error")
}
