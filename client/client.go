package client

import (
	"github.com/ajityagaty/go-kairosdb/builder"
	"github.com/ajityagaty/go-kairosdb/response"
)

type Client interface {
	// Returns a list of all metrics names.
	GetMetricNames() (*response.GetResponse, error)

	// Returns a list of all tag names.
	GetTagNames() (*response.GetResponse, error)

	// Returns a list of all tag values.
	GetTagValues() (*response.GetResponse, error)

	// Queries KairosDB using the query built using builder.
	Query(qb builder.QueryBuilder) (*response.QueryResponse, error)

	// Sends metrics from the builder to the KairosDB server.
	PushMetrics(mb builder.MetricBuilder) (*response.Response, error)

	// Deletes a metric. This is the metric and all its datapoints.
	DeleteMetric(name string) (*response.Response, error)

	// Deletes data in KairosDB using the query built by the builder.
	Delete(builder builder.QueryBuilder) (*response.Response, error)

	// Checks the health of the KairosDB Server.
	HealthCheck() (*response.Response, error)
}
