package manyfaced

import (
    "github.com/nienie/manyfaced/configuration"
    "github.com/nienie/manyfaced/source"
    "github.com/nienie/manyfaced/poll"
    "github.com/nienie/manyfaced/property"
)
//DynamicWatchedConfiguration ...
type DynamicWatchedConfiguration struct {
    *configuration.MapConfiguration
    propertyUpdater    *property.DynamicPropertyUpdater
    Source             source.WatchedConfigurationSource
}

//NewDynamicWatchedConfiguration ...
func NewDynamicWatchedConfiguration(source source.WatchedConfigurationSource) *DynamicWatchedConfiguration {
   cfg := &DynamicWatchedConfiguration{
        MapConfiguration:         configuration.NewMapConfiguration(nil),
        propertyUpdater:                    &property.DynamicPropertyUpdater{},
        Source:                             source,
    }
    currentData, err := source.GetCurrentData()
    if err != nil {
        //TODO: Add Logger
    }
    result := poll.NewFullWatchedUpdatedResult(currentData)
    cfg.UpdateConfiguration(result)
    cfg.Source.AddUpdateListener(cfg)
    return cfg
}

//UpdateConfiguration ...
func (c *DynamicWatchedConfiguration)UpdateConfiguration(result *poll.WatchedUpdateResult) {
    c.propertyUpdater.UpdateProperties(result, c)
}

//StopWatching ...
func (c *DynamicWatchedConfiguration)Close() {
    c.Source.StopWatching()
}






