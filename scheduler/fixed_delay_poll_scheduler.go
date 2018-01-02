package scheduler

import (
	"sync/atomic"
	"time"
)

const (
	//DefaultInitialDelay ...
	DefaultInitialDelay time.Duration = time.Second
	//DefaultDelay ...
	DefaultDelay time.Duration = time.Second
)

//FixedDelayPollScheduler ...
type FixedDelayPollScheduler struct {
	*BasePollScheduler
	interval     time.Duration
	initialDelay time.Duration
	stop         chan bool
	isRunning    int32
}

//NewFixedDelayPollScheduler ...
func NewFixedDelayPollScheduler(interval time.Duration, initialDelay time.Duration) *FixedDelayPollScheduler {
	pollScheduler := &FixedDelayPollScheduler{
		BasePollScheduler: &BasePollScheduler{},
		interval:          interval,
		initialDelay:      initialDelay,
		stop:              make(chan bool, 1),
		isRunning:         0,
	}
	if pollScheduler.interval <= time.Duration(0) {
		pollScheduler.interval = DefaultDelay
	}
	if pollScheduler.initialDelay <= time.Duration(0) {
		pollScheduler.initialDelay = DefaultInitialDelay
	}
	pollScheduler.BasePollScheduler.SetScheduler(pollScheduler)
	return pollScheduler
}

//Schedule ...
func (s *FixedDelayPollScheduler) Schedule(task func()) {
	if atomic.CompareAndSwapInt32(&s.isRunning, 0, 1) {
		go func() {
			defer func() {
				if r := recover(); r != nil {
					//TODO: Add Logger
				}
			}()
			initDelayed := time.After(s.initialDelay)
			select {
			case <-initDelayed:
				ticker := time.NewTicker(s.interval)
				defer ticker.Stop()
				for {
					select {
					case <-s.stop:
						s.isRunning = 0
						return
					case <-ticker.C:
						task()
					}
				}
			case <-s.stop:
				s.isRunning = 0
				return
			}
		}()
	}
}

//Stop ...
func (s *FixedDelayPollScheduler) Stop() {
	if s.isRunning == 1 {
		s.stop <- true
	}
}
