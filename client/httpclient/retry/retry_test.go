package retry

import (
	"github.com/pkg/errors"
	"github.com/saperliu/go-kairosdb/client/httpclient/backoff"
	"github.com/saperliu/go-kairosdb/client/xtime"
	"testing"
	"time"
)

func TestRetrier_Do(t *testing.T) {
	bo := backoff.NewConstantBackoff(xtime.Duration(100 * time.Millisecond))
	err := NewRetrier(bo).Do(HelloDo, 5)
	t.Log(err)
}

func HelloDo() (err error) {
	err = errors.New("retry testing")
	return
}
