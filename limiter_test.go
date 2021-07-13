package pool

import (
	"sync"
	"testing"
	"time"
)

func Test_limiter(t *testing.T) {
	var wg sync.WaitGroup

	lm := NewLimiter(2)

	fn := func() {
		defer lm.Release()()
		defer wg.Done()
		time.Sleep(time.Second)
	}

	start := time.Now()
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go fn()
	}

	wg.Wait()

	if time.Since(start) >= 3*time.Second {
		t.Fail()
	}
}
