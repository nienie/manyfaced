package scheduler

import (
    "github.com/nienie/manyfaced/configuration"
    "github.com/nienie/manyfaced/source"
    "github.com/nienie/manyfaced/property"
    "github.com/nienie/manyfaced/poll"
)

//PollScheduler ...
type PollScheduler interface {
    //StartPolling...
    StartPolling(source source.PolledConfigurationSource, config configuration.Configuration) error

    //Schedule...
    Schedule(task func())

    //Stop stop polling
    Stop()

    //AddPollListener ...
    AddPollListener(l poll.PollListener)

    //RemovePollListener ...
    RemovePollListener(l poll.PollListener)
}

//BasePollScheduler this class is responsible for scheduling the periodical polling of a configuration source and applying
//the polling result to a configuration.
type BasePollScheduler struct {
    scheduler       PollScheduler
    listeners       []poll.PollListener
    checkPoint      interface{}
    propertyUpdater property.DynamicPropertyUpdater
}

//SetScheduler ...
func (s *BasePollScheduler)SetScheduler(scheduler PollScheduler) {
    s.scheduler = scheduler
}

func (s *BasePollScheduler)initialLoad(source source.PolledConfigurationSource, config configuration.Configuration) error{
    result, err := source.Poll(true, nil)
    if err != nil {
        return err
    }
    s.checkPoint = result.GetCheckPoint()
    s.fireEvent(poll.PollEventTypeSuccess, result, nil)
    s.propertyUpdater.UpdateProperties(result, config)
    return nil
}

func (s *BasePollScheduler)fireEvent(eventType poll.PollEventType, result *poll.PolledResult, err error) {
    for _, listener := range s.listeners {
        listener.HandleEvent(eventType, result, err)
    }
}

func (s *BasePollScheduler)getPollingTask(source source.PolledConfigurationSource, config configuration.Configuration) func() {
    return func() {
        result, err := source.Poll(false, s.checkPoint)
        if err != nil {
            s.fireEvent(poll.PollEventTypeFailure, result, err)
            return
        }
        s.fireEvent(poll.PollEventTypeSuccess, result, nil)
        s.checkPoint = result.GetCheckPoint()
        s.propertyUpdater.UpdateProperties(result, config)
    }
}

//StartingPolling ...
func (s *BasePollScheduler)StartPolling(source source.PolledConfigurationSource, config configuration.Configuration) error {
    err := s.initialLoad(source, config)
    if err != nil {
        return err
    }
    task := s.getPollingTask(source, config)
    s.Schedule(task)
    return nil
}

//Stop ...
func (s *BasePollScheduler)Stop() {
    s.scheduler.Stop()
}

//Schedule ...
func (s *BasePollScheduler)Schedule(task func()) {
    s.scheduler.Schedule(task)
}

//AddPollListener ...
func (s *BasePollScheduler)AddPollListener(l poll.PollListener) {
    if l != nil {
        s.listeners = append(s.listeners, l)
    }
}

//RemovePollListener ...
func (s *BasePollScheduler)RemovePollListener(l poll.PollListener) {
    if l != nil {
        for i, listener := range s.listeners {
            if l == listener {
                s.listeners = append(s.listeners[:i], s.listeners[i+1:]...)
                break
            }
        }
    }
}