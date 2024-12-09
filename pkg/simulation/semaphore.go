package simulation

type Semaphore struct {
	channel chan struct{}
}

func NewSemaphore(limit int) *Semaphore {
	return &Semaphore{
		channel: make(chan struct{}, limit),
	}
}

func (s *Semaphore) Acquire() {
	s.channel <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.channel
}
