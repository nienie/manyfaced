package configuration

//Listener ...
type Listener interface {

	//ConfigurationChanged notifies this listener about a change on a monitored configuration object.
	ConfigurationChanged(*Event)
}
