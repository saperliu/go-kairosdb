package builder

import "encoding/json"

type MetricBuilder interface {
	// Add a new metric to the builder.
	AddMetric(name string) Metric

	// Get a list of all the metrics that are part of the builder.
	GetMetrics() []Metric

	// Encode the Metrics list into JSON.
	Build() ([]byte, error)
}

// Type that implements the MetricBuilder interface.
type mBuilder struct {
	Metrics []Metric `json:"metrics"`
}

func NewMetricBuilder() MetricBuilder {
	return &mBuilder{}
}

func (mb *mBuilder) AddMetric(name string) Metric {
	m := NewMetric(name)
	mb.Metrics = append(mb.Metrics, m)
	return m
}

func (mb *mBuilder) GetMetrics() []Metric {
	return mb.Metrics
}

func (mb *mBuilder) Build() ([]byte, error) {
	return json.Marshal(mb.Metrics)
}
