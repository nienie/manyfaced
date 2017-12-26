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
func NewDynamicFileConfiguration(configFiles []string, pollTimeout time.Duration, interval time.Duration, initialDelay time.Duration) *DynamicFileConfiguration {
    source := source.NewFileConfigurationSource(configFiles, pollTimeout)
    scheduler := scheduler.NewFixedDelayPollScheduler(interval, initialDelay)
    return &DynamicFileConfiguration{
        DynamicConfiguration:   NewDynamicConfiguration(source, scheduler),
    }
}
