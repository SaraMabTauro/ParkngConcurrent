package concurrencia

type ParkingChannel struct {
	EnterChannel chan struct{}
	ExitChannel  chan struct{}
}

func NewParkingChannel() *ParkingChannel {
	return &ParkingChannel{
		EnterChannel: make(chan struct{}),
		ExitChannel:  make(chan struct{}),
	}
}
