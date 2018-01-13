package poll

//DynamicUpdatedResult the interface represents result from a polled or watched configuration source.
//The result may be the complete content of the configuration source, or an incremental one.
type DynamicUpdatedResult interface {
	//HasChanged indicates if the configuration contents has changed.
	HasChanged() bool

	//GetComplete get complete contents from the configuration source.
	GetComplete() map[string]interface{}

	//GetAdded get the added contents from the configuration source.
	GetAdded() map[string]interface{}

	//GetChanged get the changed contents from the configuration source.
	GetChanged() map[string]interface{}

	//GetDeleted get the deleted contents from the configuraiton source.
	GetDeleted() map[string]interface{}

	//IsIncremental indicates if the result is incremental.
	IsIncremental() bool
}
