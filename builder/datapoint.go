package builder

import (
	"encoding/json"
	"errors"
)

// Represents a measurement. Stores the time when the measurement occurred and its value.
type DataPoint struct {
	timestamp int64
	value     interface{}
}

func (dp *DataPoint) Int64Value() (int64, error) {
	val, ok := dp.value.(int64)
	if !ok {
		return 0, errors.New("Not an int64 value")
	}
	return val, nil
}

func (dp *DataPoint) Float64Value() (float64, error) {
	val, ok := dp.value.(float64)
	if !ok {
		return 0, errors.New("Not a float64 value")
	}
	return val, nil
}

func (dp *DataPoint) MarshalJSON() ([]byte, error) {
	data := []interface{}{dp.timestamp, dp.value}
	return json.Marshal(data)
}

func (dp *DataPoint) UnmarshalJSON(data []byte) error {
	var arr []interface{}
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}

	var v float64
	ok := false
	if v, ok = arr[0].(float64); !ok {
		return errors.New("Invalid Timestamp type")
	}

	// Update the receiver with the values decoded.
	dp.timestamp = int64(v)
	dp.value = arr[1]

	return nil
}
