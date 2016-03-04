package aggregator

import "encoding/json"

type customAggregator struct {
	keyVal map[string]interface{}
}

func NewCustomAggregator(kv map[string]interface{}) *customAggregator {
	return &customAggregator{
		keyVal: kv,
	}
}

func (ca *customAggregator) GetName() string {
	name, ok := ca.keyVal["name"].(string)
	if !ok {
		return ""
	} else {
		return name
	}
}

func (ca *customAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(ca.keyVal)
}
