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
func NewDynamicURLConfiguration(configUrls []string, interval time.Duration,
    initialDelay time.Duration) (*DynamicURLConfiguration, error) {
    source := source.NewURLConfigurationSource(configUrls)
    scheduler := scheduler.NewFixedDelayPollScheduler(interval, initialDelay)
    dynamicConfiguration, err := NewDynamicConfiguration(source, scheduler)
    if err != nil {
        return nil, err
    }
    return &DynamicURLConfiguration{
        DynamicConfiguration:   dynamicConfiguration,
    }, nil
}

