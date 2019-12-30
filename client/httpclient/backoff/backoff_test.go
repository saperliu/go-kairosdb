package backoff

import (
	"fmt"
	"go-kairosdb/client/xtime"
	"testing"
	"time"
)

func TestExponentialBackoff_Next(t *testing.T) {
	backoff := NewExponentialBackoff(time.Second, time.Millisecond, 2.0)
	for i := 0; i <= 5; i++ {
		fmt.Printf("Interval %d \n", backoff.Next(i))
		time.Sleep(backoff.Next(i))
	}
}

func TestConstantBackoff_Next(t *testing.T) {
	backoff := NewConstantBackoff(xtime.Duration(time.Millisecond))
	for i := 0; i <= 5; i++ {
		fmt.Printf("Interval %d \n", backoff.Next(i))
		time.Sleep(backoff.Next(i))
	}
}
