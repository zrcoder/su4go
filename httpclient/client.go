package httpclient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	// HTTP methods we support
	POST    = "POST"
	GET     = "GET"
	HEAD    = "HEAD"
	PUT     = "PUT"
	DELETE  = "DELETE"
	PATCH   = "PATCH"
	OPTIONS = "OPTIONS"

	ContentTypeJson = "application/json"
	ContentTypeForm = "application/x-www-form-urlencoded"
	ContentTypeText = "text/plain"
	ContentTypeXml  = "application/xml"
	keyContentType  = "Content-Type"

	defaultTimeout = 30 * time.Second
)

type CallBackStr func(response *http.Response, body string, err error)
type CallBack func(response *http.Response, body []byte, err error)

type httpClient struct {
	url         string
	method      string
	header      map[string]string
	contentType string
	body        []byte
	transport   *http.Transport
	timeout     time.Duration
	err         error
}

func New() *httpClient {
	client := &httpClient{
		timeout: defaultTimeout,
		header:  make(map[string]string),
		transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}
	return client
}

func (c *httpClient) Get(url string) *httpClient {
	c.clearWithUrl(url)
	c.method = GET
	return c
}

func (c *httpClient) POST(url string) *httpClient {
	c.clearWithUrl(url)
	c.method = POST
	return c
}

func (c *httpClient) PUT(url string) *httpClient {
	c.clearWithUrl(url)
	c.method = PUT
	return c
}

func (c *httpClient) Head(url string) *httpClient {
	c.clearWithUrl(url)
	c.method = HEAD
	return c
}

func (c *httpClient) Delete(url string) *httpClient {
	c.clearWithUrl(url)
	c.method = DELETE
	return c
}

func (c *httpClient) Patch(url string) *httpClient {
	c.clearWithUrl(url)
	c.method = PATCH
	return c
}

func (c *httpClient) Options(url string) *httpClient {
	c.clearWithUrl(url)
	c.method = OPTIONS
	return c
}

func (c *httpClient) clearWithUrl(url string) {
	c.url = url
	c.method = ""
	c.header = make(map[string]string)
	c.contentType = ""
	c.body = nil
	c.err = nil
	c.timeout = defaultTimeout
}

func (c *httpClient) ContentType(contentType string) *httpClient {
	c.contentType = contentType
	return c
}

func (c *httpClient) Header(k, v string) *httpClient {
	if k == "" || v == "" {
		c.keepOrigionErr(errors.New("invalid header, key or value is empty"))
	} else {
		c.header[k] = v
	}
	c.header[k] = v
	return c
}

// body can be defined struct, string, map, array(or slice), and so on
func (c *httpClient) Body(body interface{}) *httpClient {
	var err error
	switch value := body.(type) {
	case string:
		c.body = []byte(value)
	case []byte:
		c.body = value
	default:
		c.body, err = json.Marshal(body)
		c.keepOrigionErr(err)
	}
	return c
}

func (c *httpClient) ConfigTls(config *tls.Config) *httpClient {
	c.transport.TLSClientConfig = config
	return c
}

func (c *httpClient) Timeout(timeout time.Duration) *httpClient {
	c.timeout = timeout
	return c
}

func (c *httpClient) keepOrigionErr(err error) {
	if c.err == nil {
		c.err = err
	}
}

func (c *httpClient) Do(callback CallBack) {
	if callback == nil {
		return
	}
	resp, body, err := c.Go()
	callback(resp, body, err)
}

func (c *httpClient) DoStr(callback CallBackStr) {
	if callback == nil {
		return
	}
	resp, body, err := c.GoStr()
	callback(resp, body, err)
}

func (c *httpClient) Go() (*http.Response, []byte, error) {
	if c.err != nil {
		return nil, nil, c.err
	}

	var (
		err  error
		resp *http.Response
		body []byte
	)

	resp, body, err = c.getResponseBytes()
	if err != nil {
		return nil, nil, err
	}

	return resp, body, nil
}

func (c *httpClient) GoStr() (*http.Response, string, error) {
	resp, body, err := c.Go()
	if err != nil {
		return nil, "", err
	}
	bodyString := string(body)
	return resp, bodyString, nil
}

func (c *httpClient) getResponseBytes() (*http.Response, []byte, error) {
	var (
		req  *http.Request
		err  error
		resp *http.Response
	)

	req, err = c.makeRequest()
	if err != nil {
		return nil, nil, fmt.Errorf("make request failed:%q", err)
	}

	client := c.makeClient()
	resp, err = client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("request failed:%q", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return resp, body, nil
}

func (c *httpClient) makeRequest() (*http.Request, error) {
	var (
		req *http.Request
		err error
	)

	req, err = http.NewRequest(c.method, c.url, bytes.NewReader(c.body))
	if err != nil {
		return nil, err
	}

	req.Header.Set(keyContentType, c.contentType)

	for k, v := range c.header {
		req.Header.Set(k, v)
	}
	return req, nil
}

func (c *httpClient) makeClient() http.Client {
	client := http.Client{
		Transport: c.transport,
		Timeout:   c.timeout,
	}
	return client
}
