package httpclient

import (
	"go-kairosdb/client/xtime"
	"testing"
	"time"
	"golang.org/x/net/context"
)

var (
	cfg    *Config
	client *HttpClient
)

func init() {
	cfg = &Config{
		Dial:            xtime.Duration(time.Second),
		Timeout:         xtime.Duration(time.Second),
		KeepAlive:       xtime.Duration(time.Second),
		BackoffInterval: xtime.Duration(time.Second),
		RetryCount:      10,
	}
	client = NewHTTPClient(cfg)
}

func TestHttpClient_Get(t *testing.T) {
	//var res interface{}
	client.SetRetryCount(5)
	res,err := client.Get(context.Background(), "https://http2.pro/api/v1", nil)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(res)

}

func TestHttpClient_Post(t *testing.T) {
	//var res interface{}
	param := make(map[string]interface{})
	res,err := client.Post(context.Background(), "https://http2.pro/api/v1", MIMEJSON, nil, param)
	if err != nil {
		t.Log(err)
	}
	t.Log(res)
}
