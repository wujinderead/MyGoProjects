package stdlib

import (
	"strconv"
	"testing"
	"time"
)

func TestFormatInt(t *testing.T) {
	var t1 time.Time = time.Now()
	var tu int64 = t1.Unix()
	t.Log(t1)
	t.Log(t1.UnixNano())
	t.Log(strconv.FormatInt(tu, 10))
	t.Log(strconv.FormatInt(14, 2))
}
