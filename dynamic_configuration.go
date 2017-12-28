package manyfaced

import (
    "github.com/nienie/manyfaced/configuration"
    "github.com/nienie/manyfaced/source"
    "github.com/nienie/manyfaced/scheduler"
)

//DynamicConfiguration a configuration that polls a PolledConfigurationSource according to the scheduler.
//The property values in this configuration will be changed dynamically at runtime. If the value changes in the configuration source.
type DynamicConfiguration struct {
    *configuration.MapConfiguration
    Source      source.PolledConfigurationSource
    Scheduler   scheduler.PollScheduler
}

//NewDynamicConfiguration ...
func NewDynamicConfiguration(source source.PolledConfigurationSource,
    scheduler scheduler.PollScheduler) (*DynamicConfiguration, error) {
    cfg := &DynamicConfiguration{
        MapConfiguration:     configuration.NewMapConfiguration(nil),
        Source:               source,
        Scheduler:            scheduler,
    }
    err := cfg.StartPolling()
    if err != nil {
        return nil, err
    }
    return cfg, nil
}

//StartPolling ...
func (c *DynamicConfiguration)StartPolling() error {
    return c.Scheduler.StartPolling(c.Source, c)
}

//StopPolling ...
func (c *DynamicConfiguration)Close() {
    if c.Scheduler != nil {
        c.Scheduler.Stop()
    }
}