package aggregator

import "github.com/ajityagaty/go-kairosdb/builder/utils"

type percentileAggregator struct {
	*samplingAggregator
	Percentile float64
}

func NewPercentileAggregator(percentile float64, value int, unit utils.TimeUnit) *percentileAggregator {
	return &percentileAggregator{
		samplingAggregator: NewSamplingAggregator("percentile", value, unit),
		Percentile:         percentile,
	}
}

func (pa *percentileAggregator) GetName() string {
	return pa.GetName()
}

func (pa *percentileAggregator) GetPercentile() float64 {
	return pa.Percentile
}
