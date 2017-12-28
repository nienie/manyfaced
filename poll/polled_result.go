package poll

//PolledResult ...
type PolledResult struct {
    *WatchedUpdateResult
    checkPoint interface{}
}

//NewFullPolledResult create a full result that represents the complete content of the configuration source.
func NewFullPolledResult(complete map[string]interface{}) *PolledResult {
    return &PolledResult{
        NewFullWatchedUpdatedResult(complete),
        nil,
    }
}

//NewIncrementalPolledResult create a result that represents incremental changes from the configuration source.
func NewIncrementalPolledResult(added, changed, deleted map[string]interface{}, checkPoint interface{}) *PolledResult {
    return &PolledResult{
        NewIncrementalWatchUpdatedResult(added, changed, deleted),
        checkPoint,
    }
}

//GetCheckPoint the check point(marker) for this poll result.
func (o *PolledResult)GetCheckPoint() interface{} {
    return o.checkPoint
}
