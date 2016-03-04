package utils

import "time"

type RelativeTime struct {
	Value int      `json:"value,omitempty"`
	Unit  TimeUnit `json:"unit,omitempty"`
}

func NewRelativeTime(value int, unit TimeUnit) *RelativeTime {
	return &RelativeTime{
		Value: value,
		Unit:  unit,
	}
}

func (rt *RelativeTime) GetValue() int {
	return rt.Value
}

func (rt *RelativeTime) GetUnit() TimeUnit {
	return rt.Unit
}

func (rt *RelativeTime) RelativeTimeTo(t time.Time) time.Time {
	var newTime time.Time

	switch rt.Unit {
	case YEARS:
		newTime = t.AddDate(-rt.Value, 0, 0)
	case MONTHS:
		newTime = t.AddDate(0, -rt.Value, 0)
	case WEEKS:
		days := rt.Value * 7
		newTime = t.AddDate(0, 0, -days)
	case DAYS:
		newTime = t.AddDate(0, 0, -rt.Value)
	case HOURS:
		newTime = t.Add(-(time.Duration(rt.Value) * time.Hour))
	case MINUTES:
		newTime = t.Add(-(time.Duration(rt.Value) * time.Minute))
	case SECONDS:
		newTime = t.Add(-(time.Duration(rt.Value) * time.Second))
	}

	return newTime
}
