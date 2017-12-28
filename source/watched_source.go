package source

import (
    "github.com/nienie/manyfaced/poll"
)

//WatchedConfigurationSource the definition of configuration source that brings dynamic changes to the configuration via watchers.
type WatchedConfigurationSource interface {

    //AddUpdateListener ...
    AddUpdateListener(poll.WatchedUpdateListener)

    //RemoveUpdateListener ...
    RemoveUpdateListener(poll.WatchedUpdateListener)

    //GetCurrentData get a snapshot of the latest configuration data.
    GetCurrentData() (map[string]interface{}, error)

    //StopWatching ...
    StopWatching()
}