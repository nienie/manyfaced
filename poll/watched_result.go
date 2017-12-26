package poll

//WatchedUpdateResult this class represents the result of a callback from the WatchedConfigurationSource.
//The result may be the complete content of the configuration source -- or an incremental one.
type WatchedUpdateResult struct {
	complete    map[string]interface{}
	added       map[string]interface{}
	changed     map[string]interface{}
	deleted     map[string]interface{}
	incremental bool
}

//NewFullWatchedUpdatedResult ...
func NewFullWatchedUpdatedResult(complete map[string]interface{}) *WatchedUpdateResult {
	result := &WatchedUpdateResult{
		complete:    make(map[string]interface{}),
		added:       make(map[string]interface{}),
		changed:     make(map[string]interface{}),
		deleted:     make(map[string]interface{}),
		incremental: false,
	}
	for key, val := range complete {
		result.complete[key] = val
	}
	return result
}

//NewIncrementalWatchUpdatedResult ...
func NewIncrementalWatchUpdatedResult(added, changed, deleted map[string]interface{}) *WatchedUpdateResult {
	result := &WatchedUpdateResult{
		complete:    make(map[string]interface{}),
		added:       make(map[string]interface{}),
		changed:     make(map[string]interface{}),
		deleted:     make(map[string]interface{}),
		incremental: true,
	}
	for key, val := range added {
		result.added[key] = val
	}
	for key, val := range changed {
		result.changed[key] = val
	}
	for key, val := range deleted {
		result.deleted[key] = val
	}
	return result
}

//HasChanged indicate whether this result has any content. If the result is incremental, this is true if
// there is any add, changed or deleted properties. If the result is complete, this is true if complete is empty.
func (o *WatchedUpdateResult) HasChanged() bool {
	if o.incremental {
		return (o.added != nil && len(o.added) > 0) ||
			(o.changed != nil && len(o.changed) > 0) ||
			(o.deleted != nil && len(o.deleted) > 0)
	}
	return o.complete != nil && len(o.complete) > 0
}

//GetComplete get the complete content from configuration source.
func (o *WatchedUpdateResult) GetComplete() map[string]interface{} {
	return o.complete
}

//GetAdded get the added properties in the configuration source as a map
func (o *WatchedUpdateResult) GetAdded() map[string]interface{} {
	return o.added
}

//GetDeleted get the deleted properties in the configuration source as a map.
func (o *WatchedUpdateResult) GetDeleted() map[string]interface{} {
	return o.deleted
}

//GetChanged get the changed properties in the configuration source as a map.
func (o *WatchedUpdateResult) GetChanged() map[string]interface{} {
	return o.changed
}

//IsIncremental whether the result is incremental. false if it is the complete content of the configuration source.
func (o *WatchedUpdateResult) IsIncremental() bool {
	return o.incremental
}
