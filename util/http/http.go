package http

import (
	"fmt"
	"go.uber.org/fx"
	"io/ioutil"
	"net/http"
	iface "red-packet/util/interface"
	"time"
)

type Client struct {
	Request *http.Request
	Timeout time.Duration
}

func (c *Client) SetTimeout(timeout time.Duration) {
	c.Timeout = timeout
}

func (c *Client) SetRequest(request *http.Request) {
	c.Request = request
}

func (c *Client) Send() (responseBody string, err error) {
	var client = &http.Client{Timeout: c.Timeout}

	response, err := client.Do(c.Request)
	if err != nil {
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		err = fmt.Errorf("response code: %d, response body: %s\n", response.StatusCode, body)
		return
	}

	return string(body), nil
}

var Module = fx.Options(
	fx.Provide(
		NewHttpClient,
	),
)

func NewHttpClient() iface.IHttpClient {
	return &Client{
		Request: nil,
		Timeout: 5 * time.Second,
	}
}
