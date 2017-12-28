package manyfaced

import (
    "time"
    "database/sql"

    "github.com/nienie/manyfaced/source"
    "github.com/nienie/manyfaced/scheduler"
)

//DynamicDBConfiguration ...
type DynamicDBConfiguration struct {
    *DynamicConfiguration
}

//NewDynamicDBConfiguration ...
func NewDynamicDBConfiguration(dbSource *sql.DB, querySQL string, keyColumnName string, valColumnName string,
    interval time.Duration, initialDelay time.Duration) (*DynamicDBConfiguration, error) {
    source := source.NewDBConfigurationSource(dbSource, querySQL, keyColumnName, valColumnName)
    scheduler := scheduler.NewFixedDelayPollScheduler(interval, initialDelay)
    dynamicConfiguration, err := NewDynamicConfiguration(source, scheduler)
    if err != nil {
        return nil, err
    }
    return &DynamicDBConfiguration{
        DynamicConfiguration:   dynamicConfiguration,
    }, nil
}
