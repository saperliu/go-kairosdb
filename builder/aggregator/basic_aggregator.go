package aggregator

type basicAggregator struct {
	Name string `json:"name,omitempty"`
}

func NewBasicAggregator(name string) *basicAggregator {
	return &basicAggregator{
		Name: name,
	}
}

func (ba *basicAggregator) GetName() string {
	return ba.Name
}
