package poll

type PollEventType int

const (
    PollEventTypeSuccess = iota
    PollEventTypeFailure
)

//PollListener the listener to be called upon when polling scheduler completes a polling.
type PollListener interface {
    //HandleEvent this method is called when the listener is invoked after a polling.
    HandleEvent(eventType PollEventType, lastResult *PolledResult, exception error)
}
