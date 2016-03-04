package builder

import "encoding/json"

// A metric contains measurements or data points. Each data point has a time
// stamp of when the measurement occurred and a value that is either a long or
// double and optionally contains tags. Tags are labels that can be added to
// better identify the metric. For example, if the measurement was done on server1
// then you might add a tag named "host" with a value of "server1". Note that a
// metric must have at least one tag.

// An interface that represents an instance of a metric.
type Metric interface {
	// Adds a TTL, expressed in seconds, to the metric.
	AddTTL(ttl int64) Metric

	// Adds a custom type of value stored in datapoint.
	AddType(t string) Metric

	// Adds a tag to the datapoint.
	AddTag(name, val string) Metric

	// Adds a datapoint to the metric. The value is of int64 type.
	AddDataPoint(timestamp int64, value interface{}) Metric

	// Returns the TLL associated with the metric.
	GetTTL() int64

	// Returns the name of the metric.
	GetName() string

	// Returns the custom type name.
	GetType() string

	// Returns all the tags/values as a map.
	GetTags() map[string]string

	// Returns an array of all the datapoints of this metric.
	GetDataPoints() []DataPoint

	// Encodes the Metric instance as a JSON array.
	Build() ([]byte, error)
}

// Type that implements the Metric interface.
type metricType struct {
	Name       string            `json:"name"`           // Name of the metric.
	Type       string            `json:"type,omitempty"` // Type of the metric being stored.
	DataPoints []DataPoint       `json:"datapoints"`     // List of DataPoints.
	Tags       map[string]string `json:"tags"`           // Map of tag names and the values associated.
	TTL        int64             `json:"ttl"`            // TTL associated with the metric.
}

func NewMetric(name string) Metric {
	return &metricType{
		Name: name,
		Tags: make(map[string]string),
	}
}

func (m *metricType) AddTTL(ttl int64) Metric {
	m.TTL = ttl
	return m
}

func (m *metricType) AddType(t string) Metric {
	m.Type = t
	return m
}

func (m *metricType) AddTag(name, val string) Metric {
	m.Tags[name] = val
	return m
}

func (m *metricType) AddDataPoint(timestamp int64, value interface{}) Metric {
	m.DataPoints = append(m.DataPoints, DataPoint{timestamp: timestamp, value: value})
	return m
}

func (m *metricType) GetName() string {
	return m.Name
}

func (m *metricType) GetType() string {
	return m.Type
}

func (m *metricType) GetTTL() int64 {
	return m.TTL
}

func (m *metricType) GetTags() map[string]string {
	return m.Tags
}

func (m *metricType) GetDataPoints() []DataPoint {
	return m.DataPoints
}

func (m *metricType) Build() ([]byte, error) {
	b, err := json.Marshal(m)
	return b, err
}
