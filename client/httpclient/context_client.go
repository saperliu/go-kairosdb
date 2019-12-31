package httpclient

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/saperliu/go-kairosdb/client/httpclient/backoff"
	"github.com/saperliu/go-kairosdb/client/httpclient/retry"
	"github.com/saperliu/go-kairosdb/client/xtime"
	"golang.org/x/net/http2"
	"io"
	"net"
	xhttp "net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	minRead               = 16 * 1024 // 16kb
	defaultRetryCount int = 0
)

type Config struct {
	Dial            xtime.Duration
	Timeout         xtime.Duration
	KeepAlive       xtime.Duration
	MaxConns        int
	MaxIdle         int
	BackoffInterval xtime.Duration // Interval is second
	RetryCount      int
}

type HttpClient struct {
	conf       *Config
	client     *xhttp.Client
	dialer     *net.Dialer
	transport  *xhttp.Transport
	retryCount int
	retrier    retry.Retriable
}

// NewHTTPClient returns a new instance of httpClient
func NewHTTPClient(c *Config) *HttpClient {
	dialer := &net.Dialer{
		Timeout:   time.Duration(c.Dial),
		KeepAlive: time.Duration(c.KeepAlive),
	}
	transport := &xhttp.Transport{
		DialContext:         dialer.DialContext,
		MaxConnsPerHost:     c.MaxConns,
		MaxIdleConnsPerHost: c.MaxIdle,
		IdleConnTimeout:     time.Duration(c.KeepAlive),
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}
	_ = http2.ConfigureTransport(transport)
	bo := backoff.NewConstantBackoff(c.BackoffInterval)
	return &HttpClient{
		conf: c,
		client: &xhttp.Client{
			Transport: transport,
		},
		retryCount: defaultRetryCount,
		retrier:    retry.NewRetrier(bo),
	}
}

// SetRetryCount sets the retry count for the httpClient
func (c *HttpClient) SetRetryCount(count int) {
	c.retryCount = count
}

// SetRetryCount sets the retry count for the httpClient
func (c *HttpClient) SetRetrier(retrier retry.Retriable) {
	c.retrier = retrier
}

// Get makes a HTTP GET request to provided URL with context passed in
func (c *HttpClient) Get(ctx context.Context, url string, headers xhttp.Header) (response *xhttp.Response, err error) {
	request, err := xhttp.NewRequest(xhttp.MethodGet, url, nil)
	if err != nil {
		return nil,errors.Wrap(err, "GET - request creation failed")
	}

	request.Header = headers

	return c.Do(ctx, request)
}

// Post makes a HTTP POST request to provided URL with context passed in
func (c *HttpClient) Post(ctx context.Context, url, contentType string, headers xhttp.Header, param []byte) (response *xhttp.Response, err error) {
	request, err := xhttp.NewRequest(xhttp.MethodPost, url, bytes.NewReader(param))
	if err != nil {
		return nil,errors.Wrap(err, "POST - request creation failed")
	}
	if headers == nil {
		headers = make(xhttp.Header)
	}
	headers.Set("Content-Type", contentType)
	request.Header = headers

	//logs.Info("------Post request------ %v   %v ",request ,param)
	return c.Do(ctx, request)
}

// Put makes a HTTP PUT request to provided URL with context passed in
func (c *HttpClient) Put(ctx context.Context, url, contentType string, headers xhttp.Header, param []byte) (response *xhttp.Response, err error) {
	request, err := xhttp.NewRequest(xhttp.MethodPut, url, bytes.NewReader(param))
	if err != nil {
		return nil,errors.Wrap(err, "PUT - request creation failed")
	}

	if headers == nil {
		headers = make(xhttp.Header)
	}
	headers.Set("Content-Type", contentType)
	request.Header = headers

	return c.Do(ctx, request)
}

// Patch makes a HTTP PATCH request to provided URL with context passed in
func (c *HttpClient) Patch(ctx context.Context, url, contentType string, headers xhttp.Header, param []byte) (response *xhttp.Response, err error) {
	request, err := xhttp.NewRequest(xhttp.MethodPatch, url, bytes.NewReader(param))
	if err != nil {
		return nil,errors.Wrap(err, "PATCH - request creation failed")
	}

	if headers == nil {
		headers = make(xhttp.Header)
	}
	headers.Set("Content-Type", contentType)
	request.Header = headers

	return c.Do(ctx, request)
}

// Delete makes a HTTP DELETE request to provided URL with context passed in
func (c *HttpClient) Delete(ctx context.Context, url, contentType string, headers xhttp.Header, param interface{}) (response *xhttp.Response, err error) {
	request, err := xhttp.NewRequest(xhttp.MethodDelete, url, nil)
	if err != nil {
		return nil,errors.Wrap(err, "DELETE - request creation failed")
	}

	if headers == nil {
		headers = make(xhttp.Header)
	}
	headers.Set("Content-Type", contentType)
	request.Header = headers

	return c.Do(ctx, request)
}

// Do makes an HTTP request with the native `http.Do` interface and context passed in
func (c *HttpClient) Do(ctx context.Context, req *xhttp.Request) (response *xhttp.Response, err error) {
	for i := 0; i <= c.retryCount; i++ {
		if response, err = c.request(ctx, req); err != nil {
			backoffTime := c.retrier.NextInterval(i)
			time.Sleep(backoffTime)
			continue
		}
		break
	}
	return response, err
}

func (c *HttpClient) request(ctx context.Context, req *xhttp.Request) (response *xhttp.Response, err error) {
	var (
		//response *xhttp.Response
		//bs       []byte
		cancel func()
	)
	ctx, cancel = context.WithTimeout(ctx, time.Duration(c.conf.Timeout))
	defer cancel()
	response, err = c.client.Do(req.WithContext(ctx))
	if err != nil {
		select {
		case <-ctx.Done():
			err = ctx.Err()
		}
		return
	}
	return response, err
	//defer response.Body.Close()
	//if response.StatusCode >= xhttp.StatusInternalServerError {
	//	err = errors.Wrap(err, "")
	//	return
	//}
	//if bs, err = readAll(response.Body, minRead); err != nil {
	//	return
	//}
	//err = json.Unmarshal(bs, &res)
	//return
}

func reqBody(contentType string, param []byte) (body io.Reader) {
	//var err error
	if contentType == MIMEPOSTForm {
		//enc, ok := param.(string)
	//	if ok {
	//		body = strings.NewReader(enc)
	//	}
	}
	if contentType == MIMEJSON {
		//buff := new(bytes.Buffer)
		//err = json.NewEncoder(buff).Encode(param)
		//if err != nil {
		//	return
		//}
		//body = buff
		fmt.Printf(" body  %s",string(param))
		body = bytes.NewReader(param)
	}
	return
}

func readAll(r io.Reader, capacity int64) (b []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0, capacity))
	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(r)
	return buf.Bytes(), err
}
