package scheduler

import (
    "time"
    "sync"
)

const (
    //DefaultInitialDelay ...
    DefaultInitialDelay time.Duration = time.Second
    //DefaultDelay ...
    DefaultDelay        time.Duration = time.Second
)

//FixedDelayPollScheduler ...
type FixedDelayPollScheduler struct {
    sync.Mutex
    *BasePollScheduler
    interval     time.Duration
    initialDelay time.Duration
    stop         chan bool
    isRunning    bool
}

//NewFixedDelayPollScheduler ...
func NewFixedDelayPollScheduler(interval time.Duration, initialDelay time.Duration) *FixedDelayPollScheduler {
    pollScheduler := &FixedDelayPollScheduler{
        Mutex:              sync.Mutex{},
        BasePollScheduler:  &BasePollScheduler{},
        interval:           interval,
        initialDelay:       initialDelay,
        stop:               make(chan bool),
        isRunning:          false,
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
func (s *FixedDelayPollScheduler)Schedule(task func()) {
    if s.isRunning {
        return
    }
    s.Lock()
    defer s.Unlock()
    if s.isRunning == false {
        s.isRunning = true
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
                defer func(){ ticker.Stop()}()
                for {
                    select {
                    case <-s.stop:
                        s.isRunning = false
                        return
                    case <-ticker.C:
                        task()
                    }
                }
            case <-s.stop:
                s.isRunning = false
                return
            }
        }()
    }
}

//Stop ...
func (s *FixedDelayPollScheduler)Stop() {
    s.stop <- true
}
