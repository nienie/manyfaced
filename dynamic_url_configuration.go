package manyfaced

import (
    "time"

    "github.com/nienie/manyfaced/source"
    "github.com/nienie/manyfaced/scheduler"
)

//DynamicURLConfiguration ...
type DynamicURLConfiguration struct {
    *DynamicConfiguration
}

//NewDynamicURLConfiguration ...
func NewDynamicURLConfiguration(configUrls []string, interval time.Duration, initialDelay time.Duration) *DynamicURLConfiguration {
    source := source.NewURLConfigurationSource(configUrls)
    scheduler := scheduler.NewFixedDelayPollScheduler(interval, initialDelay)
    return &DynamicURLConfiguration{
        DynamicConfiguration:   NewDynamicConfiguration(source, scheduler),
    }
}

