package pool

type Limiter interface {
	Release() func()
}

func NewLimiter(size int) Limiter {
	return &limiter{ch: make(chan bool, size)}
}

type limiter struct {
	ch chan bool
}

func (l *limiter) Release() func() {
	l.ch <- true
	return func() {
		<-l.ch
	}
}
