package builder

type Aggregator interface {
	// Gets the name of the aggregation being used.
	GetName() string
}
