package manyfaced

import (
    "time"

    "github.com/nienie/manyfaced/source"
    "github.com/nienie/manyfaced/scheduler"
)

//DynamicFileConfiguration ...
type DynamicFileConfiguration struct {
    *DynamicConfiguration
}

//NewDynamicFileConfiguration ...
func NewDynamicFileConfiguration(configFiles []string, pollTimeout time.Duration,
    interval time.Duration, initialDelay time.Duration) (*DynamicFileConfiguration, error) {
    source := source.NewFileConfigurationSource(configFiles, pollTimeout)
    scheduler := scheduler.NewFixedDelayPollScheduler(interval, initialDelay)
    dynamicConfiguration, err := NewDynamicConfiguration(source, scheduler)
    if err != nil {
        return nil, err
    }
    return &DynamicFileConfiguration{
        DynamicConfiguration:   dynamicConfiguration,
    }, nil
}
