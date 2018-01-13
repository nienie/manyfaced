package poll

//WatchedUpdateListener ...
type WatchedUpdateListener interface {
	//UpdateConfiguration updates the configuration.
	UpdateConfiguration(*WatchedUpdateResult)
}
