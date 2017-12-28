package poll

//EventType ...
type EventType int

const (
    //EventTypeSuccess ...
    EventTypeSuccess = iota
    //EventTypeFailure ...
    EventTypeFailure
)

//Listener the listener to be called upon when polling scheduler completes a polling.
type Listener interface {
    //HandleEvent this method is called when the listener is invoked after a polling.
    HandleEvent(eventType EventType, lastResult *PolledResult, exception error)
}
