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
