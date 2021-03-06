// Copyright 2016 Ajit Yagaty
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"encoding/json"
	"fmt"
	"github.com/saperliu/go-kairosdb/builder"
	"github.com/saperliu/go-kairosdb/client/httpclient"
	"github.com/saperliu/go-kairosdb/client/xtime"
	"github.com/saperliu/go-kairosdb/response"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	api_version      = "/api/v1"
	datapoints_ep    = api_version + "/datapoints"
	deldatapoints_ep = api_version + "/datapoints/delete"
	query_ep         = api_version + "/datapoints/query"
	health_ep        = api_version + "/health/check"
	delmetric_ep     = api_version + "/metric/"
	metricnames_ep   = api_version + "/metricnames"
	tagnames_ep      = api_version + "/tagnames"
	tagvalues_ep     = api_version + "/tagvalues"
	version_ep       = api_version + "/version"
)

// This is the type that implements the Client interface.
type httpClient struct {
	serverAddress string
	Client        *httpclient.HttpClient
}

func (hc *httpClient) NewClient(serverAddress string) Client {
	cfg := &httpclient.Config{
		Dial:            xtime.Duration(time.Second),
		Timeout:         xtime.Duration(time.Second),
		KeepAlive:       xtime.Duration(time.Second),
		BackoffInterval: xtime.Duration(time.Second),
		RetryCount:      3,
	}
	hc.Client = httpclient.NewHTTPClient(cfg)
	return &httpClient{
		serverAddress: serverAddress,
		Client:        hc.Client,
	}
}

func NewHttpClient(serverAddress string) Client {
	cfg := &httpclient.Config{
		Dial:            xtime.Duration(time.Second),
		Timeout:         xtime.Duration(time.Second),
		KeepAlive:       xtime.Duration(time.Second),
		BackoffInterval: xtime.Duration(time.Second),
		RetryCount:      3,
	}
	client := httpclient.NewHTTPClient(cfg)
	return &httpClient{
		serverAddress: serverAddress,
		Client:        client,
	}
}

// Returns a list of all metrics names.
func (hc *httpClient) GetMetricNames() (*response.GetResponse, error) {
	return hc.get(hc.serverAddress + metricnames_ep)
}

// Returns a list of all tag names.
func (hc *httpClient) GetTagNames() (*response.GetResponse, error) {
	return hc.get(hc.serverAddress + tagnames_ep)
}

// Returns a list of all tag values.
func (hc *httpClient) GetTagValues() (*response.GetResponse, error) {
	return hc.get(hc.serverAddress + tagvalues_ep)
}

// Queries KairosDB using the query built using builder.
func (hc *httpClient) Query(qb builder.QueryBuilder) (*response.QueryResponse, error) {
	// Get the JSON representation of the query.
	data, err := qb.Build()
	if err != nil {
		return nil, err
	}

	return hc.postQuery(hc.serverAddress+query_ep, data)
}

// Sends metrics from the builder to the KairosDB server.
func (hc *httpClient) PushMetrics(mb builder.MetricBuilder) (*response.Response, error) {
	data, err := mb.Build()
	if err != nil {
		return nil, err
	}

	return hc.postData(hc.serverAddress+datapoints_ep, data)
}

// Deletes a metric. This is the metric and all its datapoints.
func (hc *httpClient) DeleteMetric(name string) (*response.Response, error) {
	return hc.delete(hc.serverAddress + delmetric_ep + name)
}

// Deletes data in KairosDB using the query built by the builder.
func (hc *httpClient) Delete(qb builder.QueryBuilder) (*response.Response, error) {
	data, err := qb.Build()
	if err != nil {
		return nil, err
	}

	return hc.postData(hc.serverAddress+deldatapoints_ep, data)
}

// Checks the health of the KairosDB Server.
func (hc *httpClient) HealthCheck() (*response.Response, error) {
	resp, err := hc.sendRequest(hc.serverAddress+health_ep, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	r := &response.Response{}
	r.SetStatusCode(resp.StatusCode)
	return r, nil
}

func (hc *httpClient) sendRequest(url, method string) (*http.Response, error) {
	//req, err := http.NewRequest(method, url, nil)
	//if err != nil {
	//	return nil, err
	//}
	//req.Header.Add("accept", "application/json")
	//cli := &http.Client{}
	//
	//return cli.Do(req)
	headers := make(http.Header)
	headers.Set("accept", "application/json")
	headers.Set("Content-Type", "application/json")
	return hc.Client.Get(context.Background(), url, headers)
}

func (hc *httpClient) httpRespToResponse(httpResp *http.Response) (*response.Response, error) {
	defer httpResp.Body.Close()
	resp := &response.Response{}
	resp.SetStatusCode(httpResp.StatusCode)

	if httpResp.StatusCode != http.StatusNoContent {
		// If the request has failed, then read the response body.
		//defer httpResp.Body.Close()
		contents, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return nil, err
		} else {
			// Unmarshal the contents into Response object.
			fmt.Printf(" kairosdb  response : %v ", contents)
			//logs.Info("-----contents  -------   %v  ",string(contents))
			err = json.Unmarshal(contents, resp)
			if err != nil {
				return nil, err
			}
		}
	}

	return resp, nil
}

func (hc *httpClient) httpRespToQueryResponse(httpResp *http.Response) (*response.QueryResponse, error) {
	// Read the HTTP response body.
	defer httpResp.Body.Close()
	contents, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	qr := response.NewQueryResponse(httpResp.StatusCode)

	// Unmarshal the contents into QueryResponse object.
	err = json.Unmarshal(contents, qr)
	if err != nil {
		return nil, err
	}

	return qr, nil
}

func (hc *httpClient) get(url string) (*response.GetResponse, error) {
	resp, err := hc.sendRequest(url, "GET")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	} else {
		gr := response.NewGetResponse(resp.StatusCode)

		err = json.Unmarshal(contents, gr)
		if err != nil {
			return nil, err
		}

		return gr, nil
	}
}

func (hc *httpClient) postData(url string, data []byte) (*response.Response, error) {
	//resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Printf("-----postData  finish ------- %v    %v    %v  ", url, err, resp)
	//return hc.httpRespToResponse(resp)
	resp, err := hc.Client.Post(context.Background(), url, "application/json", nil, data)
	if err != nil {
		return nil, err
	}
	fmt.Printf("-----postData  finish ------- %v    %v    %v  ", url, err, resp)
	return hc.httpRespToResponse(resp)
}

func (hc *httpClient) postQuery(url string, data []byte) (*response.QueryResponse, error) {
	//resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	//if err != nil {
	//	return nil, err
	//}
	headers := make(http.Header)
	headers.Set("accept", "application/json")
	headers.Set("Content-Type", "application/json")
	resp, err := hc.Client.Post(context.Background(), url, "application/json", headers, data)
	if err != nil {
		return nil, err
	}
	return hc.httpRespToQueryResponse(resp)
}

func (hc *httpClient) delete(url string) (*response.Response, error) {
	resp, err := hc.sendRequest(url, "DELETE")
	if err != nil {
		return nil, err
	}

	return hc.httpRespToResponse(resp)
}
