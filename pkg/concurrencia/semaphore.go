package semaphore

type Semaphore struct {
	permits chan struct{}
}

func NewSemaphore(maxPermits int) *Semaphore {
	return &Semaphore{
		permits: make(chan struct{}, maxPermits),
	}
}

func (s *Semaphore) Acquire() {
	s.permits <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.permits
}

func (s *Semaphore) TryAcquire() bool {
	select {
	case s.permits <- struct{}{}:
		return true
	default:
		return false
	}
}

func (s *Semaphore) Available() int {
	return cap(s.permits) - len(s.permits)
}
