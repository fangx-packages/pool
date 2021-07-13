package pool

import (
	"testing"
	"time"
)

func Test_limiter(t *testing.T) {
	lm := NewLimiter(2)

	fn := func() {
		defer lm.Release()()
		time.Sleep(time.Second)
	}

	start := time.Now()

	for i := 0; i < 3; i++ {
		go fn()
	}

	if time.Since(start) >= 3*time.Second {
		t.Fail()
	}
}
