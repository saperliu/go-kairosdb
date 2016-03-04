package aggregator

import "github.com/ajityagaty/go-kairosdb/builder/utils"

type rateAggregator struct {
	*basicAggregator
	unit utils.TimeUnit
}

func NewRateAggregator(unit utils.TimeUnit) *rateAggregator {
	return &rateAggregator{
		basicAggregator: NewBasicAggregator("rate"),
		unit:            unit,
	}
}

func (ra *rateAggregator) GetUnit() utils.TimeUnit {
	return ra.unit
}
