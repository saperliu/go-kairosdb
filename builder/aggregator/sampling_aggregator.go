package aggregator

import "github.com/ajityagaty/go-kairosdb/builder/utils"

type sampling struct {
	Value int            `json:"value,omitempty"`
	Unit  utils.TimeUnit `json:"unit,omitempty"`
}

type samplingAggregator struct {
	*basicAggregator
	//Name           string
	AlignStartTime bool     `json:"align_start_time,omitempty"`
	AlignSampling  bool     `json:"align_sampling,omitempty"`
	StartTime      int64    `json:"start_time,omitempty"`
	Sample         sampling `json:"sampling,omitempty"`
}

func NewSamplingAggregator(name string, value int, unit utils.TimeUnit) *samplingAggregator {
	return &samplingAggregator{
		basicAggregator: NewBasicAggregator(name),
		//name:            name,
		Sample: sampling{
			Value: value,
			Unit:  unit,
		},
	}
}

func (sa *samplingAggregator) GetValue() int {
	return sa.Sample.Value
}

func (sa *samplingAggregator) GetUnit() utils.TimeUnit {
	return sa.Sample.Unit
}

// Alignment based on the sampling size. For example if your sample size is either milliseconds,
// seconds, minutes or hours then the start of the range will always be at the top
// of the hour.  The effect of setting this to true is that your data will
// take the same shape when graphed as you refresh the data.
//
// Only one alignment type can be used.
func (sa *samplingAggregator) WithSamplingAlignment() *samplingAggregator {
	sa.AlignSampling = true
	return sa
}

// Alignment based on the aggregation range rather than the value of the first
// data point within that range.
//
// Only one alignment type can be used.
func (sa *samplingAggregator) WithStartTimeAlignment(startTime int64) *samplingAggregator {
	sa.AlignStartTime = true
	sa.StartTime = startTime
	return sa
}

func (sa *samplingAggregator) IsAlignSampling() bool {
	return sa.AlignSampling
}

func (sa *samplingAggregator) IsAlignStartTime() bool {
	return sa.AlignStartTime
}

func (sa *samplingAggregator) GetAlignStartTime() int64 {
	return sa.StartTime
}

/*
func (sa *samplingAggregator) GetName() string {
	return "tmp"
}
*/
